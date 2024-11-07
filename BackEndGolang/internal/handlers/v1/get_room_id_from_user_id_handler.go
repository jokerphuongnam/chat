package handlers

import (
	"chat-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *Handler) GetRoomIDFromUserID(c *gin.Context) {
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
	anotherUserIDStr := c.Param("user_id")
	anotherUserID, err := uuid.Parse(anotherUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid room ID" + err.Error(),
		})
		return
	}

	roomID, err := handler.Database.GetRoomIDFromUserID(userID, anotherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get room ID" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "Room ID retrieved successfully",
		Data:    roomID,
	})
}
