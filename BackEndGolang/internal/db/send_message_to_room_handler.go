package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/ent/member"
	"chat-backend/internal/ent/message"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Content     string              `json:"content"`
	TypeMessage message.TypeMessage `json:"type_message"`
	SendTime    uint64              `json:"send_time"`
	Sender      string              `json:"sender_id"`
	Room        uuid.UUID           `json:"room_id"`
}

func SendMessageToRoomHandler(client *ent.Client, from, to uuid.UUID, message string, messageType message.TypeMessage) (MessageResponse, error) {
	// Check if the user is part of the room.
	member, err := client.Member.Query().
		Where(
			member.And(member.UserIDEQ(from), member.RoomID(to)),
		).
		Only(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("user is not part of the room: %w", err)
	}

	// Create a new message.
	newMessage, err := client.Message.Create().
		SetIDUserSend(from).
		SetIDRoom(member.RoomID).
		SetContent(message).
		SetDateSend(uint64(context.Background().Value("timestamp").(uint64))).
		SetTypeMessage(messageType).
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to create message: %w", err)
	}

	return MessageResponse{
		Content:     newMessage.Content,
		TypeMessage: newMessage.TypeMessage,
		SendTime:    newMessage.DateSend,
		Sender:      from.String(),
		Room:        member.RoomID,
	}, nil
}
