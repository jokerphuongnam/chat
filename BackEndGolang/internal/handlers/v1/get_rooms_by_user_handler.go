package handlers

import (
	database "chat-backend/internal/db"
	"chat-backend/internal/ent"
	"chat-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoomsByUserHandler(c *gin.Context, dbClient *ent.Client, secretKey string) {
	// Get user ID from the header.
	userID, err := utils.GetUserIdFromHeader(c, secretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	roomDetails, err := database.GetRoomsByUserHandler(dbClient, userID)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
            Code:    http.StatusInternalServerError,
            Message: err.Error(),
        })
        return
	}
	
	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
        Message: "success",
        Data:    roomDetails,
	})
}
