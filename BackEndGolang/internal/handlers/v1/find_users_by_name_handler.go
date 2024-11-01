package handlers

import (
	database "chat-backend/internal/db"
	"chat-backend/internal/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsersByNameHandler(c *gin.Context, dbClient *ent.Client) {
	// Get name from the query parameters.
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "name parameter is required",
		})
		return
	}

	// Call FindUsersByName function
	users, err := database.FindUsersByNameHandler(dbClient, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "failed to find users",
		})
		return
	}

	// Success response with users data
	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    users,
	})
}
