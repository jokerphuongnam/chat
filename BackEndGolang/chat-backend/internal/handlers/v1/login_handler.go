package handlers

import (
	"chat-backend/internal/logs"
	"chat-backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (handler *Handler) LoginHandler(c *gin.Context) {
	logs.Log.Debug("Login start...")

	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	response, err := handler.Database.LoginHandler(request.Username, request.Password)
	if err != nil {
		logs.Log.Errorf("database error: %v", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Invalid username or password",
		})
		return
	}

	// Generate a JWT token for the user and update the authorize record with the token
	jwtToken, err := utils.GenerateJWT(response.ID.String(), handler.Cache.SecretKey)
	if err != nil {
		logs.Log.Error("failed to generate JWT token: ", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate JWT token" + err.Error(),
		})
		return
	}

	// Cache the JWT token
	err = handler.Cache.CacheJWTToken(response.ID.String(), jwtToken)
	if err != nil {
		logs.Log.Error("failed to cache JWT token: ", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to cache JWT token: %v", err),
		})
		return
	}

	response.JwtToken = jwtToken

	logs.Log.Debug("Login success")
	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "login successfully",
		Data:    response,
	})
}
