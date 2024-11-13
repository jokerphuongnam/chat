package handlers

import (
	"chat-backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=4,max=16"`
	Name     string `json:"name" binding:"required"`
}

func (handler *Handler) RegisterHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	// Create a new user in the database
	response, err := handler.Database.RegisterHandler(req.Username, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Generate a JWT token for the user and update the authorize record with the token
	jwtToken, err := utils.GenerateJWT(response.ID.String(), handler.Cache.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate JWT token" + err.Error(),
		})
		return
	}

	// Cache the JWT token
	err = handler.Cache.CacheJWTToken(response.ID.String(), jwtToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to cache JWT token: %v", err),
		})
		return
	}

	response.JwtToken = jwtToken

	c.JSON(http.StatusCreated, ResponseMessage{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
		Data:    response,
	})
}
