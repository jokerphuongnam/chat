package handlers

import (
	"chat-backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (handler *Handler) ChatHandler(c *gin.Context) {
	// Get jwt token from headers
	jwtToken, err := utils.GetJWTTokenFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseMessage{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized: Missing or invalid Bearer token",
		})
		return
	}

	// Get user ID from the header.
	userID, err := handler.Cache.GetUserID(jwtToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Upgrade WebSocket connection.
	conn, err := (*handler.ChatService).UpgradeConnection(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	defer conn.Close()

	// Connect user to the chat service.
	if err := (*handler.ChatService).ConnectUser(userID, jwtToken, conn); err != nil {
		conn.Close()
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	defer (*handler.ChatService).DisconnectUser(jwtToken)

	// Listen for incoming messages from the client and broadcast them to other connected users.
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Read error from user %s - token %v: %v\n", userID, jwtToken, err), http.StatusInternalServerError)
			break
		}
		if len(message) == 0 {
			(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("received empty message from user %s - token %v with message type: %v\n", userID, jwtToken, messageType), http.StatusInternalServerError)
			break
		}
		if messageType == websocket.TextMessage {
			var req *sendMessageRequest
			if err := json.Unmarshal(message, &req); err != nil {
				(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Read error from user %s: %v\n", userID, err), http.StatusInternalServerError)
				break
			}

			handler.sendMessageHandler(conn, jwtToken, userID, *req)
		}
	}
}

type sendMessageRequest struct {
	To          string `json:"to"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (handler *Handler) sendMessageHandler(conn *websocket.Conn, jwtToken string, userID uuid.UUID, req sendMessageRequest) {
	toID, err := utils.StringToUUID(req.To)
	if err != nil {
		(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Failed to parse user ID: %v", err), http.StatusInternalServerError)
		return
	}

	toUserId, roomUsers, err := handler.Database.CheckUserInRoomHandler(toID)
	if err != nil {
		(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("User not found in the room: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the message type.
	messageType, err := utils.ParseTypeMessage(req.MessageType)
	if err != nil {
		(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Failed to parse message type: %v", err), http.StatusInternalServerError)
		return
	}
	if toUserId != nil {
		newMessage, err := handler.Database.SendMessageToNewUserHandler(userID, toID, req.Message, messageType)
		if err != nil {
			(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Failed to save message to the database: %v", err), http.StatusInternalServerError)
			return
		}

		// Send the message to the user.
		(*handler.ChatService).SendMessage(
			conn,
			jwtToken,
			userID,
			req.Message,
			req.MessageType,
			newMessage.SendTime,
			newMessage.Room,
			[]uuid.UUID{*toUserId})
		return
	} else if roomUsers != nil {
		newMessage, err := handler.Database.SendMessageToRoomHandler(userID, roomUsers.RoomId, req.Message, messageType)
		if err != nil {
			(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("Failed to save message to the database: %v", err), http.StatusInternalServerError)
			return
		}

		// Send the message to all the users in the room.
		(*handler.ChatService).SendMessage(
			conn,
			jwtToken,
			userID,
			req.Message,
			req.MessageType,
			newMessage.SendTime,
			newMessage.Room,
			roomUsers.UserIds,
		)
		return
	}

	(*handler.ChatService).SendErrorMessageByConnection(conn, fmt.Sprintf("user not found in the room:"), http.StatusInternalServerError)
}
