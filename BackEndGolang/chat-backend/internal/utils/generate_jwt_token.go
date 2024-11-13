package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func jwtKey(secretKey string) []byte {
	return []byte(secretKey)
}

func GenerateJWT(userID, secretKey string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expiration time

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	return token.SignedString(jwtKey(secretKey))
}

func ParseJWT(tokenString, secretKey string) (uuid.UUID, error) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey(secretKey), nil
	})
	if err!= nil ||!token.Valid {
        return uuid.UUID{}, fmt.Errorf("invalid token: %v", err)
    }
	userID, err := uuid.Parse(claims.Subject)
	if err!= nil {
        return uuid.UUID{}, fmt.Errorf("invalid UUID format: %v", err)
    }
	return userID, nil
}