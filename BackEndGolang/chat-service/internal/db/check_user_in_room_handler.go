package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type RoomUsers struct {
	RoomId  uuid.UUID
	UserIds []uuid.UUID
}

func (db *Database) CheckUserInRoomHandler(id uuid.UUID) (*uuid.UUID, *RoomUsers, error) {
	room, err := db.Client.Room.Get(context.Background(), id)
	if err == nil {
		userIDs := make([]uuid.UUID, 0)

		members, err := room.QueryMembers().WithUsers().All(context.Background())
		if err != nil {
			return nil, nil, err
		}

		for _, member := range members {
			userIDs = append(userIDs, member.UserID)
		}

		return nil, &RoomUsers{RoomId: room.ID, UserIds: userIDs}, nil
	}

	user, err := db.Client.User.Get(context.Background(), id)
	if err == nil {
		return &user.ID, nil, nil
	}

	return nil, nil, fmt.Errorf("could not retrieve user rooms: %w", err)
}
