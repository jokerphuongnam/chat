package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	webSocketURL    = "ws://localhost:3031/v1/ws"
	loginURL        = "http://localhost:3031/v1/login"
	connectionCount = 10000
)

// Client represents an authenticated client with a WebSocket connection.
type Client struct {
	Token string
	Conn  *websocket.Conn
}

// LoginRequest represents the login request payload.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the structure of the login response.
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

// MessagePayload represents the structure of the message to send.
type MessagePayload struct {
	To          string `json:"to"`
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

// authenticateUser performs a login to obtain a JWT token.
func authenticateUser(username, password string) (string, error) {
	// Prepare the login request payload
	loginRequest := LoginRequest{
		Username: username,
		Password: password,
	}

	// Convert login request to JSON
	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login request: %v", err)
	}

	// Send POST request to login URL
	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("login request failed: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal login response: %v", err)
	}

	// Return the token from the response data
	return loginResponse.Data.Token, nil
}

// connectWebSocket establishes a WebSocket connection with a token in the header.
func connectWebSocket(token string) (*websocket.Conn, error) {
	header := http.Header{}
	header.Add("Authorization", "Bearer "+token)
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, header)
	return conn, err
}

// handleMessages simulates sending a structured message on WebSocket.
func handleMessages(client *Client, wg *sync.WaitGroup) {
	// Prepare the message payload
	messagePayload := MessagePayload{
		To:          "7c13d441-5168-40aa-a740-36c0edc2a9f4",
		MessageType: "text",
		Message:     "dfsdfbigdfgdfgdgdsfghghdfgcvx mvxefd",
	}

	// Convert message to JSON

	message, err := json.Marshal(messagePayload)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return
	}

	// Send the message
	go func() {
		for i := 0; i < 10; i++ {
			err = client.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		defer wg.Done()
	}()
}

// RunLoadTest executes the load test with the specified number of connections.
func RunLoadTest() {
	var wg sync.WaitGroup

	// Creating and connecting clients
	var tokens = make([]string, 0)
	for i := 0; i < connectionCount; i++ {
		token, err := authenticateUser("phuongnam1234", "phuongnam")
		if err != nil {
			log.Fatalf("Error during login: %v", err)
		}
		tokens = append(tokens, token)
		log.Printf("Connect to token: %s", token)
	}

	var isEnd = false

	for _, token := range tokens {
		conn, err := connectWebSocket(token)
		if err != nil {
			log.Fatalf("WebSocket connection error: %v", err)
		}

		// Read response (if any)
		go func() {
			for {
				_, response, err := conn.ReadMessage()
				if err != nil {
					log.Printf("Error reading message: %v", err)
					return
				}
				log.Printf("Received: %s", response)
				if isEnd {
                    conn.Close()
                    return
                }
			}
		}()
		client := &Client{Token: token, Conn: conn}
		wg.Add(1)
		go handleMessages(client, &wg)
	}

	// Wait for all clients to finish
	wg.Wait()
	log.Println("Load test completed.")
	isEnd = true
}


func main() {
    RunLoadTest()
}