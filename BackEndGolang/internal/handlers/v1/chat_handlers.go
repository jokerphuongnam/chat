package handlers

import (
	"chat-backend/internal/ent"
	"chat-backend/internal/services"
	"chat-backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context, cs services.ChatService, dbClient *ent.Client, secretKey string) {
	// Get user ID from the header.
	userID, err :=  utils.GetUserIdFromHeader(c, secretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Upgrade WebSocket connection.
	conn, err := cs.UpgradeConnection(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}

	// Connect user to the chat service.
	if err := cs.ConnectUser(userID, conn); err != nil {
		conn.Close()
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	defer cs.DisconnectUser(userID)

	// Listen for incoming messages from the client and broadcast them to other connected users.
	for {
		_, message, err := conn.ReadMessage()
		if err!= nil {
            fmt.Printf("Read error from user %s: %v\n", userID, err)
            break
        }
		fmt.Printf("Received message from %s: %s\n", userID, message)
	}
}
