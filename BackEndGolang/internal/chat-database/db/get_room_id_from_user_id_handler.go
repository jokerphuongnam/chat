package database

import (
	"context"
	"fmt"

	"chat-database/ent"
	"chat-database/ent/member"
	"chat-database/ent/room"

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
			return &anotherUserID, nil
		}
		return nil, fmt.Errorf("failed to retrieve roomID: %w", err)
	}

	if roomEntity.Edges.RoomInfo == nil {
		return &roomEntity.ID, nil
	}

	return nil, nil
}
