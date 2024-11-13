package handlers

import (
	"chat-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetUserInfoHandler(c *gin.Context) {
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

	user, err := handler.Database.GetUserInfoHandler(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    user,
	})
}
