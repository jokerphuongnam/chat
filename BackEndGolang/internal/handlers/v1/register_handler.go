package handlers

import (
	database "chat-backend/internal/db"
	"chat-backend/internal/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=4,max=16"`
	Name     string `json:"name" binding:"required"`
}

func RegisterHandler(c *gin.Context, dbClient *ent.Client, secretKey string) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	// Create a new user in the database
	response, err := database.RegisterHandler(dbClient, secretKey, req.Username, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ResponseMessage{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
		Data:    response,
	})
}
