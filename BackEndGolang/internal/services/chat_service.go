package services

import (
	"chat-backend/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatService interface {
	// Upgrade a HTTP request to a WebSocket connection
	UpgradeConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)

	// Initialize a WebSocket connection for a specific user by user ID
	ConnectUser(userID uuid.UUID, conn *websocket.Conn) error

	// Disconnect a user and remove their WebSocket connection from the pool
	DisconnectUser(userID uuid.UUID) error

	// Send a direct message from users to another
	SendMessage(from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, getToIds func() []uuid.UUID)

	// Broadcast a message to all active connections (optional, depending on requirements)
	BroadcastMessage(message string) error
}

type chatService struct {
	mu          sync.Mutex
	connections map[uuid.UUID]*websocket.Conn
	upgrader    websocket.Upgrader
}

func NewChatService() ChatService {
	return &chatService{
		connections: make(map[uuid.UUID]*websocket.Conn),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (cs *chatService) UpgradeConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	conn, err := cs.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (cs *chatService) ConnectUser(userID uuid.UUID, conn *websocket.Conn) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, exists := cs.connections[userID]; exists {
		return errors.New("user already connected")
	}
	cs.connections[userID] = conn
	return nil
}

func (cs *chatService) DisconnectUser(userID uuid.UUID) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conn, exists := cs.connections[userID]
	if !exists {
		return errors.New("user not connected")
	}

	conn.Close()
	delete(cs.connections, userID)
	return nil
}

type message struct {
	MessageType string    `json:"message_type"`
	Content     string    `json:"message"`
	SenderID    uuid.UUID `json:"sender_id"`
	SendAt      uint64    `json:"sent_at"`
}

func (cs *chatService) SendMessage(from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, getToIds func() []uuid.UUID) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	toIds := getToIds()
	connections := cs.connections

	for userID, conn := range connections {
		if userID != from && utils.ArrayExists(userID, toIds) {
			msg, err := json.Marshal(
				message{
					MessageType: messageType,
					Content:     content,
					SenderID:    from,
					SendAt:      sendAt,
				},
			)
			if err != nil {
				fmt.Printf("Error marshalling message: %v\n", err)
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Printf("Error sending message to user %s: %v\n", userID, err)
			}
		}
	}
}

func (cs *chatService) BroadcastMessage(message string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for userID, conn := range cs.connections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Printf("Error broadcasting to user %s: %v\n", userID, err)
		}
	}
	return nil
}
