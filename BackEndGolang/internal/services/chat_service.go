package services

import (
	"chat-backend/internal/logs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatService interface {
	// Upgrade a HTTP request to a WebSocket connection
	UpgradeConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)

	// Initialize a WebSocket connection for a specific user by user ID
	ConnectUser(userID uuid.UUID, token string, conn *websocket.Conn) error

	// Disconnect a user and remove their WebSocket connection from the pool
	DisconnectUser(token string) error

	// Send a direct message from users to another
	SendMessage(fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID)

	// Send error message
	SendErrorMessage(fromToken string, from uuid.UUID, message string, code int)

	SendErrorMessageByConnection(conn *websocket.Conn, message string, code int)

	// Broadcast a message to all active connections (optional, depending on requirements)
	BroadcastMessage(message string) error
}

type chatService struct {
	connections    map[uuid.UUID]map[string]*websocket.Conn
	upgrader       websocket.Upgrader
	messageChan    chan *messageChan
	disconnectChan chan string
	secretKey      string
}

type messageChan struct {
	message   responseMessage
	roomID    uuid.UUID
	toIDs     []uuid.UUID
	from      uuid.UUID
	fromToken string
}

type responseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewChatService(secretKey string) ChatService {
	cs := &chatService{
		secretKey:      secretKey,
		connections:    make(map[uuid.UUID]map[string]*websocket.Conn),
		messageChan:    make(chan *messageChan),
		disconnectChan: make(chan string),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	go cs.handleChannels()
	return cs
}

func (cs *chatService) UpgradeConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	conn, err := cs.upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusBadRequest)
		return nil, err
	}
	return conn, nil
}

func (cs *chatService) ConnectUser(userID uuid.UUID, token string, conn *websocket.Conn) error {
	logs.Log.Debugf("connecting user %v...\n", userID)
	// If the user is already connected, check for existing tokens
	if tokenConnections, exists := cs.connections[userID]; exists {
		// Check if the token is already in the user's connection map
		if _, tokenExists := tokenConnections[token]; tokenExists {
			return errors.New("user already connected with the given token")
		}
	}

	if _, exists := cs.connections[userID]; !exists {
		cs.connections[userID] = make(map[string]*websocket.Conn)
	}

	cs.connections[userID][token] = conn
	logs.Log.Debugf("user %v with token %v connected\n", userID, token)
	return nil
}

func (cs *chatService) DisconnectUser(token string) error {
	cs.disconnectChan <- token
	return nil
}

type message struct {
	MessageType string    `json:"message_type"`
	Content     string    `json:"message"`
	SenderID    uuid.UUID `json:"sender_id"`
	SendAt      uint64    `json:"sent_at"`
}

func (cs *chatService) SendMessage(fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID) {
	logs.Log.Debugf("Sending message from %v to room %v to users %v...\n", from, roomID, toIds)
	cs.messageChan <- &messageChan{
		message: responseMessage{
			Code:    http.StatusOK,
			Message: "Send Message successfully",
			Data: message{
				MessageType: messageType,
				Content:     content,
				SenderID:    from,
				SendAt:      sendAt,
			},
		},
		roomID:    roomID,
		toIDs:     toIds,
		from:      from,
		fromToken: fromToken,
	}
}

func (cs *chatService) SendErrorMessage(fromToken string, from uuid.UUID, message string, code int) {
	conn := cs.connections[from][fromToken]
	if conn == nil {
		logs.Log.Debugf("User %v with token %v not found, skipping message broadcast\n", from, fromToken)
		return
	}
	cs.SendErrorMessageByConnection(conn, message, code)
}

func (cs *chatService) SendErrorMessageByConnection(conn *websocket.Conn, message string, code int) {
	data, err := json.Marshal(responseMessage{
		Code:    code,
		Message: message,
	})
	if err != nil {
		logs.Log.Errorf("Error marshalling message: %v\n", err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, data)
}

func (cs *chatService) BroadcastMessage(message string) error {
	return nil
}

func (cs *chatService) handleChannels() {
	for {
		select {
		case msg := <-cs.messageChan:
			cs.processMessage(msg)
		case token := <-cs.disconnectChan:
			cs.processDisconnect(token)
		}
	}
}

func (cs *chatService) processMessage(msg *messageChan) {
	for _, userID := range msg.toIDs {
		if tokenConns, exists := cs.connections[userID]; exists {
			logs.Log.Debugf("to %v, userID: %v", msg.toIDs, userID)
			for token, conn := range tokenConns {
				logs.Log.Debugf("from %v, userID: %v, token %v", msg.from, userID, token)
				if token == msg.fromToken {
					data, err := json.Marshal(responseMessage{
						Code:    http.StatusOK,
						Message: "Send Message successfully",
					})
					if err != nil {
						(*cs).SendErrorMessageByConnection(conn, fmt.Sprintf("Error marshalling message: %v\n", err), http.StatusInternalServerError)
						return
					}
					if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
						(*cs).SendErrorMessageByConnection(conn, fmt.Sprintf("Error sending message to user %s (token %s): %v\n", userID, token, err), http.StatusInternalServerError)
					}
				} else {
					data, err := json.Marshal(msg.message)
					if err != nil {
						(*cs).SendErrorMessageByConnection(conn, fmt.Sprintf("Error marshalling message: %v\n", err), http.StatusInternalServerError)
						return
					}
					if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
						(*cs).SendErrorMessageByConnection(conn, fmt.Sprintf("Error sending message to user %s (token %s): %v\n", userID, token, err), http.StatusInternalServerError)
					}
				}
			}
		}
	}

	logs.Log.Debugf("Send message success from %v to room %v to users %v...\n", msg.from, msg.roomID, msg.toIDs)
}

func (cs *chatService) processDisconnect(token string) {
	var userID uuid.UUID

	for id, userConns := range cs.connections {
		if conn, exists := userConns[token]; exists {
			userID = id

			conn.Close()

			delete(userConns, token)

			logs.Log.Debugf("Disconnected token %v for user %v\n", token, userID)

			if len(userConns) == 0 {
				delete(cs.connections, userID)
				logs.Log.Debugf("User %v has no more connections, removed from pool\n", userID)
			}
			return
		}
	}

	logs.Log.Warningf("Token %v not found in any active connection pool\n", token)
}
