package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/ent/member"
	"chat-backend/internal/ent/message"
	"chat-backend/internal/ent/user"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type GetMessagesResponse struct {
	ID            uuid.UUID `json:"id"`              // Message ID
	SenderID      uuid.UUID `json:"sender_id"`       // Sender ID
	Time          uint64    `json:"time"`            // Timestamp of the message
	Content       string    `json:"content"`         // Message content
	ContentType   string    `json:"content_type"`    // Type of message (text, image, etc.)
	Nickname      string    `json:"nickname"`        // Nickname of sender
	AvatarURL     string    `json:"avatar_url"`      // Avatar URL of sender
	Color         string    `json:"color"`           // Color of room
	IsCurrentUser bool      `json:"is_current_user"` // Is current user
}

func (db *Database) GetMessagesByRoomIdHandler(roomID uuid.UUID, currentID uuid.UUID) ([]GetMessagesResponse, error) {
	messages, err := db.Client.Message.
		Query().
		Where(message.IDRoomEQ(roomID)).
		WithUsers(func(q *ent.UserQuery) {
			q.Select(user.FieldID, user.FieldName, user.FieldAvatarURL)
		}).
		WithRooms(func(q *ent.RoomQuery) {
			q.WithMembers(func(q *ent.MemberQuery) {
				q.Select(member.FieldNickName, member.FieldUserID)
			})
		}).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}

	var responseMessages []GetMessagesResponse

	for _, msg := range messages {
		nickname := msg.Edges.Rooms.Edges.Members[0].NickName
		if nickname == "" {
			nickname = msg.Edges.Users.Name
		}

		senderId := msg.Edges.Users.ID
		responseMessages = append(responseMessages, GetMessagesResponse{
			ID:            msg.ID,
			SenderID:      senderId,
			Time:          msg.DateSend,
			Content:       msg.Content,
			ContentType:   msg.TypeMessage.String(),
			Nickname:      nickname,
			AvatarURL:     msg.Edges.Users.AvatarURL,
			Color:         msg.Edges.Rooms.Color,
			IsCurrentUser: senderId == currentID,
		})
	}

	return responseMessages, nil
}
