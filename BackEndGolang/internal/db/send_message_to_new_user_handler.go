package database

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/ent/message"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func SendMessageToNewUserHandler(client *ent.Client, from, to uuid.UUID, message string, messageType message.TypeMessage) (MessageResponse, error) {
	// Step 1: Create a new room
	newRoom, err := client.Room.Create().
		SetID(to).
		SetColor("#0000FF").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("error creating new room: %v", err)
	}

	// Step 2: Create Member records for both users
	_, err = client.Member.Create().
		SetUserID(from).
		SetRoomID(newRoom.ID).
		SetRole("ADMIN").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to create Member for 'from' user: %w", err)
	}

	_, err = client.Member.Create().
		SetUserID(to).
		SetRoomID(newRoom.ID).
		SetRole("USER").
		Save(context.Background())
	if err != nil {
		return MessageResponse{}, fmt.Errorf("failed to create Member for 'to' user: %w", err)
	}

	// Step 4: Create the message in the newly created room
	newMessage, err := client.Message.Create().
		SetIDUserSend(from).
		SetIDRoom(newRoom.ID).
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
		Room: newRoom.ID,
	}, nil
}
