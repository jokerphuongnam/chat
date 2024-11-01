package handlers

import (
	database "chat-backend/internal/db"
	"chat-backend/internal/ent"
	"chat-backend/internal/services"
	"chat-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SendMessageRequest struct {
	To          string `json:"to"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func SendMessageHandler(c *gin.Context, dbClient *ent.Client, cs services.ChatService, secretKey string) {
	// Get user ID from the header.
	userID, err := utils.GetUserIdFromHeader(c, secretKey)
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

	toUserId, roomUsers, err := database.CheckUserInRoomHandler(dbClient, toID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Parse the message type.
	messageType, err := utils.ParseTypeMessage(req.MessageType)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if toUserId == nil {
		// Save the message to the database.
		newMessage, err := database.SendMessageToNewUserHandler(dbClient, userID, toID, req.Message, messageType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseMessage{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		// Send the message to the user.
		cs.SendMessage(userID, req.Message, req.MessageType, newMessage.SendTime, newMessage.Room, func() []uuid.UUID {
			return []uuid.UUID{*toUserId}
		})

		c.JSON(http.StatusCreated, ResponseMessage{
			Code:    http.StatusCreated,
			Message: "Message sent successfully",
			Data:    newMessage,
		})

		return
	} else if roomUsers != nil {
		// Save the message to the database.
		newMessage, err := database.SendMessageToRoomHandler(dbClient, userID, roomUsers.RoomId, req.Message, messageType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseMessage{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		// Send the message to all the users in the room.
		cs.SendMessage(userID, req.Message, req.MessageType, newMessage.SendTime, newMessage.Room, func() []uuid.UUID {
			return roomUsers.UserIds
		})

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
