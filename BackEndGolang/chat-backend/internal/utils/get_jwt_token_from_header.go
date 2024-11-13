package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetJWTTokenFromHeader(c *gin.Context) (string, error) {
	// Get JWT token from header
	token := c.GetHeader("Authorization")
	if token == "" {
		return "", errors.New("unauthorized: Missing Authorization header")
	}

	// Extract token from Bearer format: "Bearer <token>"
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))
	if token == "" {
		return "", errors.New("unauthorized: Invalid Authorization format")
	}
	return token, nil
}
