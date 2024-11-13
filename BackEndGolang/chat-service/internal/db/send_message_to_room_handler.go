package database

import (
	"chat-service/internal/ent/member"
	"chat-service/internal/ent/message"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Content     string              `json:"content"`
	TypeMessage message.TypeMessage `json:"type_message"`
	SendTime    uint64              `json:"send_time"`
	Sender      string              `json:"sender_id"`
	Room        uuid.UUID           `json:"room_id"`
}

func (db *Database) SendMessageToRoomHandler(from, to uuid.UUID, message string, messageType message.TypeMessage) (MessageResponse, error) {
	// Start transaction
	tx, err := db.Client.Tx(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // Rollback on panic
			panic(p)          // Re-throw the panic after rollback
		} else if err != nil {
			_ = tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	// Check if the user is part of the room.
	member, err := db.Client.Member.Query().
		Where(
			member.And(member.UserIDEQ(from), member.RoomID(to)),
		).
		Only(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("user is not part of the room: %w", err)
	}

	// Create a new message.
	newMessage, err := db.Client.Message.Create().
		SetIDUserSend(from).
		SetIDRoom(member.RoomID).
		SetContent(message).
		SetDateSend(uint64(time.Now().UnixMilli())).
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
