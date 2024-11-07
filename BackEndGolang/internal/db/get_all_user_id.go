package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (db *Database) GetAllUsersId() ([]uuid.UUID, error) {
	users, err := db.Client.User.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error fetching users from database: %v", err)
	}

	var userIDs []uuid.UUID
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	return userIDs, nil
}
