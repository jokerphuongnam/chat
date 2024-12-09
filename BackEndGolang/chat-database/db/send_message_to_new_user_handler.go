package database

import (
	"chat-database/ent"
	"chat-database/ent/member"
	"chat-database/ent/message"
	"chat-database/ent/room"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (db *Database) SendMessageToNewUserHandler(from, to uuid.UUID, message string, messageType message.TypeMessage) (MessageResponse, error) {
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

	roomEntity, err := db.Client.Room.Query().
		Where(
			room.HasMembersWith(member.UserID(from)),
			room.HasMembersWith(member.UserID(to)),
		).
		WithRoomInfo().
		Only(context.Background())

	fmt.Printf("Room exist %v", roomEntity)
	if err == nil {
		if roomEntity.Edges.RoomInfo == nil {
			return MessageResponse{}, fmt.Errorf("user already inbox")
		}
	} else if !ent.IsNotFound(err) {
		return MessageResponse{}, fmt.Errorf("failed to query existing room: %w", err)
	}

	// Step 1: Create a new room
	newRoom, err := db.Client.Room.Create().
		SetColor("#0000FF").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("error creating new room: %v", err)
	}

	// Step 2: Create Member records for both users
	_, err = db.Client.Member.Create().
		SetUserID(from).
		SetRoomID(newRoom.ID).
		SetRole("ADMIN").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to create Member for 'from' user: %w", err)
	}

	_, err = db.Client.Member.Create().
		SetUserID(to).
		SetRoomID(newRoom.ID).
		SetRole("ADMIN").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to create Member for 'to' user: %w", err)
	}

	// Step 4: Create the message in the newly created room
	newMessage, err := db.Client.Message.Create().
		SetIDUserSend(from).
		SetIDRoom(newRoom.ID).
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
		Room:        newRoom.ID,
	}, nil
}
