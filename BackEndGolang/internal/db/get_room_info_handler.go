package database

import (
	"chat-backend/internal/ent/room"
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type GetRoomInfoResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	RoomImage *string   `json:"room_image"`
	RoomName  string    `json:"room_name"`
	Color     string    `json:"color"`
	IsInbox   bool      `json:"is_inbox"`
}

func (db *Database) GetRoomInfoHandler(userID uuid.UUID, roomId uuid.UUID) (*GetRoomInfoResponse, error) {
	roomEntity, err := db.Client.Room.Query().Where(room.ID(roomId)).WithMembers().WithRoomInfo().Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get room info")
	}
	roomInfo := GetRoomInfoResponse{
		ID:      roomEntity.ID,
		Color:   roomEntity.Color,
		IsInbox: roomEntity.Edges.RoomInfo == nil,
	}

	if roomEntity.Edges.RoomInfo == nil {
		for _, member := range roomEntity.Edges.Members {
			if member.UserID != userID {
				roomInfo.RoomImage = &member.Edges.Users.AvatarURL
				if member.NickName == "" {
					roomInfo.RoomName = member.Edges.Users.Name
				} else {
					roomInfo.RoomName = member.NickName
				}

				break
			}
		}
	} else {
		if roomEntity.Edges.RoomInfo.RoomImageURL == "" {
			roomInfo.RoomImage = nil
		} else {
			roomInfo.RoomImage = &roomEntity.Edges.RoomInfo.RoomImageURL
		}

		if roomEntity.Edges.RoomInfo.Name == "" {
			var memberNames []string
			for _, member := range roomEntity.Edges.Members {
				memberNames = append(memberNames, member.Edges.Users.Name)
			}
			roomInfo.RoomName = strings.Join(memberNames, ", ")
		} else {
			roomInfo.RoomName = roomEntity.Edges.RoomInfo.Name
		}
	}

	return &roomInfo, nil
}
