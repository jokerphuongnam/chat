package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/utils"
	"context"
	"fmt"
)

type RegisterResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

func RegisterHandler(client *ent.Client, secretKey, username, password, name string) (*RegisterResponse, error) {
	// Validate the username and password format
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password %w", err)
	}

	// Create a new username/password and authorize
	up, err := client.UsernamePassword.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create username/password: %w", err)
	}

	// Create a new authorize
	au, err := client.Authorize.Create().
		SetToken(up.ID.String()).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create authorize: %w", err)
	}

	// Create a new user
	user, err := client.User.Create().
		SetName(name).
		SetIDAuthorize(au.ID).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate a JWT token for the user and update the authorize record with the token
	jwtToken, err := utils.GenerateJWT(user.ID.String(), secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %w", err)
	}

	// Save the updated JWT token in the authorize record
	auUpdate, err := client.Authorize.UpdateOneID(au.ID).SetJwtToken(jwtToken).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to update JWT token: %w", err)
	}

	return &RegisterResponse{
		Name:     user.Name,
		Username: up.Username,
		Token:    auUpdate.JwtToken,
	}, nil
}
