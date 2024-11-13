package handlers

import (
	"chat-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SendMessageRequest struct {
	To          string `json:"to"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (handler *Handler) SendMessageHandler(c *gin.Context) {
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

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	toID, err := utils.StringToUUID(req.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	toUserId, roomUsers, err := handler.Database.CheckUserInRoomHandler(toID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Parse the message type.
	messageType, err := utils.ParseTypeMessage(req.MessageType)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if toUserId != nil {
		// Save the message to the database.
		newMessage, err := handler.Database.SendMessageToNewUserHandler(userID, toID, req.Message, messageType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseMessage{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		// Send the message to the user.
		(*handler.ChatService).SendMessage(nil, jwtToken, userID, req.Message, req.MessageType, newMessage.SendTime, newMessage.Room, []uuid.UUID{*toUserId})

		c.JSON(http.StatusCreated, ResponseMessage{
			Code:    http.StatusCreated,
			Message: "Message sent successfully",
			Data:    newMessage,
		})
		return
	} else if roomUsers != nil {
		// Save the message to the database.
		newMessage, err := handler.Database.SendMessageToRoomHandler(userID, roomUsers.RoomId, req.Message, messageType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseMessage{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		// Send the message to all the users in the room.
		(*handler.ChatService).SendMessage(nil, jwtToken, userID, req.Message, req.MessageType, newMessage.SendTime, newMessage.Room, roomUsers.UserIds)

		c.JSON(http.StatusCreated, ResponseMessage{
			Code:    http.StatusCreated,
			Message: "Message sent successfully",
			Data:    newMessage,
		})
		return
	}

	c.JSON(http.StatusBadRequest, ResponseMessage{
		Code:    http.StatusBadRequest,
		Message: "User not found",
	})
}
