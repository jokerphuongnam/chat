package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIdFromHeader(c *gin.Context, secretKey string) (uuid.UUID, error) {
	// Get JWT token from header.
	token := c.GetHeader("Authorization")
	if token == "" {
		return uuid.Nil, errors.New("Unauthorized")
	}

	// Parse JWT token to get user ID.
	userID, err := ParseJWT(token, secretKey)
	if err != nil {
		return uuid.Nil, errors.New("Invalid user ID")
	}

	return userID, nil
}