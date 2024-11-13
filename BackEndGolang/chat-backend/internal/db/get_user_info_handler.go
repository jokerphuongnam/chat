package database

import (
	"chat-backend/internal/ent/user"
	"context"

	"github.com/google/uuid"
)

type GetUserInfoResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
}

func (db *Database) GetUserInfoHandler(userID uuid.UUID) (*GetUserInfoResponse, error) {
	user, err := db.Client.User.Query().Where(user.IDEQ(userID)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return &GetUserInfoResponse{
		ID:        user.ID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}
