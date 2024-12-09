package services

import (
	"chat-logs/logs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

const (
	natsTopic = "chat.messages.topic"
)

type ChatService interface {
	// Upgrade a HTTP request to a WebSocket connection
	UpgradeConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)

	// Initialize a WebSocket connection for a specific user by user ID
	ConnectUser(userID uuid.UUID, token string, conn *websocket.Conn) error

	// Disconnect a user and remove their WebSocket connection from the pool
	DisconnectUser(token string)

	// Send a direct message from users to another
	SendMessage(conn *websocket.Conn, fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID)

	// Send error message
	SendErrorMessage(fromToken string, from uuid.UUID, message string, code int)

	// Send error message by WebSocket connection
	SendErrorMessageByConnection(conn *websocket.Conn, message string, code int)

	// Broadcast a message to all active connections (optional, depending on requirements)
	BroadcastMessage(message string) error

	// Deinitialize all connections
	DeinitChatService()
}

type chatService struct {
	connections    map[uuid.UUID]map[string]*websocket.Conn
	upgrader       websocket.Upgrader
	messageChan    chan *messageChan
	disconnectChan chan string
	secretKey      string
	natsConn       *nats.Conn
	mu             sync.Mutex
	subscription   *nats.Subscription
}

type messageChan struct {
	Message   message     `json:"message"`
	RoomID    uuid.UUID   `json:"room_id"`
	ToIDs     []uuid.UUID `json:"to_ids"`
	From      uuid.UUID   `json:"from"`
	FromToken string      `json:"from_token"`
}

type responseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewChatService(secretKey string, natsConn *nats.Conn) ChatService {
	cs := &chatService{
		secretKey:      secretKey,
		natsConn:       natsConn,
		mu:             sync.Mutex{},
		subscription:   nil,
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
	go cs.receivedMessageFromNats()
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
	cs.mu.Lock()
	defer cs.mu.Unlock()

	// If the user is already connected, check for existing tokens
	if tokenListeners, exists := cs.connections[userID]; exists {
		// Check if the token is already in the user's connection map
		if _, tokenExists := tokenListeners[token]; tokenExists {
			return errors.New("user already connected with the given token")
		}
	}

	if _, exists := cs.connections[userID]; !exists {
		cs.connections[userID] = make(map[string]*websocket.Conn)
	}

	cs.connections[userID][token] = conn

	logs.Log.Infof("user %v with token %v connected\n", userID, token)
	return nil
}

func (cs *chatService) DisconnectUser(token string) {
	cs.disconnectChan <- token
}

type message struct {
	MessageType string    `json:"message_type"`
	Content     string    `json:"message"`
	SenderID    uuid.UUID `json:"sender_id"`
	SendAt      uint64    `json:"sent_at"`
}

func (cs *chatService) SendMessage(conn *websocket.Conn, fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID) {
	cs.sendMessageToNats(fromToken, from, content, messageType, sendAt, roomID, toIds)
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

func (cs *chatService) DeinitChatService() {
	cs.subscription.Unsubscribe()
	cs.natsConn.Close()
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
	currentHost, _ := os.Hostname()
	logs.Log.Infof(fmt.Sprintf("ProcessMessage websocket send at host: %v from %v", currentHost, msg.From))

	for _, userID := range msg.ToIDs {
		if tokenConns, exists := cs.connections[userID]; exists {
			for token, conn := range tokenConns {
				logs.Log.Infof("websocket send at host: %v from %v, userID: %v, token %v", currentHost, msg.From, userID, token)
				if token == msg.FromToken {
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
					data, err := json.Marshal(responseMessage{
						Code:    http.StatusOK,
						Message: "Send Message successfully",
						Data:    msg.Message,
					})
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

	logs.Log.Infof("Send message success from %v to room %v to users %v...\n", msg.From, msg.RoomID, msg.ToIDs)
}

func (cs *chatService) processDisconnect(token string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for userID, userConns := range cs.connections {
		if conn, exists := userConns[token]; exists {
			conn.Close()

			delete(cs.connections[userID], token)

			logs.Log.Infof("Disconnected token %v for user %v\n", token, userID)

			if len(userConns) == 0 {
				cs.mu.Lock()
				delete(cs.connections, userID)
			}
		}
	}
}

func (cs *chatService) sendMessageGoroutine(fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID) {
	cs.messageChan <- &messageChan{
		Message: message{
			MessageType: messageType,
			Content:     content,
			SenderID:    from,
			SendAt:      sendAt,
		},
		RoomID:    roomID,
		ToIDs:     toIds,
		From:      from,
		FromToken: fromToken,
	}
}

func (cs *chatService) sendMessageToNats(fromToken string, from uuid.UUID, content, messageType string, sendAt uint64, roomID uuid.UUID, toIds []uuid.UUID) {
	mc := messageChan{
		Message: message{
			MessageType: messageType,
			Content:     content,
			SenderID:    from,
			SendAt:      sendAt,
		},
		RoomID:    roomID,
		ToIDs:     toIds,
		From:      from,
		FromToken: fromToken,
	}

	data, err := mcToBytes(mc)
	if err != nil {
		cs.SendErrorMessage(fromToken, from, fmt.Sprintf("Error marshalling message to bytes: %v\n", err), http.StatusInternalServerError)
		return
	}

	currentHost, _ := os.Hostname()
	logs.Log.Infof("Befor sending message to via NATS: %v", currentHost)
	if err := cs.natsConn.Publish(natsTopic, data); err != nil {
		cs.SendErrorMessage(fromToken, from, fmt.Sprintf("Error marshalling message to bytes: %v\n", err), http.StatusInternalServerError)
	}
}

func (cs *chatService) receivedMessageFromNats() error {
	sub, err := cs.natsConn.Subscribe(natsTopic, func(msg *nats.Msg) {
		var response messageChan

		if err := json.Unmarshal(msg.Data, &response); err != nil {
			logs.Log.Errorf("Error unmarshalling received message: %v\n", err)
		}

		logs.Log.Infof("received message from NATS: %+v\n", response)
		cs.sendMessageGoroutine(response.FromToken, response.From, response.Message.Content, response.Message.MessageType, response.Message.SendAt, response.RoomID, response.ToIDs)
	})

	logs.Log.Infof("get subscription from nats server")
	cs.subscription = sub

	return err
}

func mcToBytes(mc messageChan) ([]byte, error) {
	data, err := json.Marshal(mc)
	if err != nil {
		return nil, fmt.Errorf("error marshalling message: %v", err)
	}
	return data, nil
}
