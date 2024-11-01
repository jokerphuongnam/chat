package handlers

import (
	database "chat-backend/internal/db"
	"chat-backend/internal/ent"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context, dbClient *ent.Client, secretKey string) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	response, err := database.LoginHandler(dbClient, secretKey, request.Username, request.Password)
	if err != nil {
		fmt.Printf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "login successfully",
		Data:    response,
	})
}
