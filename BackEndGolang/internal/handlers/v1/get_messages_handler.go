package handlers

import (
	"chat-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *Handler) GetMessagesHandler(c *gin.Context) {
	// Get jwt token from headers
	jwtToken, err := utils.GetJWTTokenFromHeader(c)
	if err!= nil {
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

	// Parse room ID from the URL parameter.
	roomIDStr := c.Param("room_id")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid room ID",
		})
		return
	}

	// Retrieve messages for the specified room
	messages, err := handler.Database.GetMessagesByRoomIdHandler(roomID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch messages",
		})
		return
	}

	// Return the messages in a ResponseMessage
	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "Messages retrieved successfully",
		Data:    messages,
	})
}
