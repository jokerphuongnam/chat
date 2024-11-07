package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/ent/member"
	"chat-backend/internal/ent/room"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (db *Database) GetRoomIDFromUserID(userID uuid.UUID, anotherUserID uuid.UUID) (*uuid.UUID, error) {
	roomEntity, err := db.Client.Room.
		Query().
		Where(
			// Ensure the room has both members
			room.HasMembersWith(
				member.UserID(userID),
			),
			room.HasMembersWith(
				member.UserID(anotherUserID),
			),
		).
		Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve roomID: %w", err)
	}

	if roomEntity.Edges.RoomInfo == nil {
		return nil, fmt.Errorf("room does not have room info")
	}

	return nil, nil
}
