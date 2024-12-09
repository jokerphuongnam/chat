package database

import (
	"chat-database/ent/authorize"
	"chat-database/ent/user"
	"chat-database/ent/usernamepassword"
	"chat-database/internal/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type LoginResponse struct {
	ID        uuid.UUID `json:"-"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
	JwtToken  string    `json:"token"`
}

func (db *Database) LoginHandler(username, password string) (*LoginResponse, error) {
	// Validate the username and password format
	usernamePassword, err := db.Client.UsernamePassword.Query().Where(usernamepassword.UsernameEQ(username)).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid username: %w", err)
	}
	if !utils.CheckPasswordHash(usernamePassword.Password, password) {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	authorizeRecord, err := db.Client.Authorize.Query().Where(authorize.TokenEQ(usernamePassword.ID.String())).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid Authorize: %w", err)
	}

	user, err := db.Client.User.Query().Where(user.IDAuthorizeEQ(authorizeRecord.ID)).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid user: %w", err)
	}

	return &LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}
