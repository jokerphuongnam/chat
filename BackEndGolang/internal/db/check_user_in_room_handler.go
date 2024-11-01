package database

import (
	"chat-backend/internal/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type RoomUsers struct {
	RoomId  uuid.UUID
	UserIds []uuid.UUID
}

func CheckUserInRoomHandler(client *ent.Client, id uuid.UUID) (*uuid.UUID, *RoomUsers, error) {
	user, err := client.User.Get(context.Background(), id)
	if err == nil {
		return &user.ID, nil, nil
	}

	room, err := client.Room.Get(context.Background(), id)
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

	return nil, nil, fmt.Errorf("could not retrieve user rooms: %w", err)
}
