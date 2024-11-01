package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/ent/authorize"
	"chat-backend/internal/ent/user"
	"chat-backend/internal/ent/usernamepassword"
	"chat-backend/internal/utils"
	"context"
	"fmt"
)

type LoginResponse struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	JwtToken  string `json:"token"`
}

func LoginHandler(dbClient *ent.Client, secretKey, username, password string) (*LoginResponse, error) {
	// Validate the username and password format
	usernamePassword, err := dbClient.UsernamePassword.Query().Where(usernamepassword.UsernameEQ(username)).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid username: %w", err)
	}
	if !utils.CheckPasswordHash(usernamePassword.Password, password) {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	authorizeRecord, err := dbClient.Authorize.Query().Where(authorize.TokenEQ(usernamePassword.ID.String())).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid Authorize: %w", err)
	}

	user, err := dbClient.User.Query().Where(user.IDAuthorizeEQ(authorizeRecord.ID)).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid user: %w", err)
	}

	jwtToken, err := utils.GenerateJWT(user.ID.String(), secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %v", err)
	}

	auUpdate, err := dbClient.Authorize.UpdateOneID(authorizeRecord.ID).SetJwtToken(jwtToken).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to update JWT token: %v", err)
	}

	return &LoginResponse{
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
		JwtToken:  auUpdate.JwtToken,
	}, nil
}
