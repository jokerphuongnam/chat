package cmd

import (
	"chat-cache/cache"
	"chat-config/config"
	database "chat-database/db"
	"chat-service/internal/handlers/v1"
	"chat-service/internal/services"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func Execute(config config.AppConfig) *gin.Engine {
	nc, err := nats.Connect(config.Nats.Addr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer nc.Close()

	r := gin.Default()
	databaseHandler, err := database.GetClient(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer databaseHandler.Client.Close()

	chatService := services.NewChatService(config.Server.SecretKey, nc)
	chatService.DeinitChatService()

	handler := &handlers.Handler{
		Database:    databaseHandler,
		ChatService: &chatService,
		Cache: &cache.Cache{
			Client:    cache.NewRedisClient(config.Cache.Addr),
			SecretKey: config.Server.SecretKey,
		},
	}

	rootCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the server",
		Run: func(cmd *cobra.Command, arg []string) {
			r.Use(cors.New(cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With", "X-Custom-Header", "Access-Control-Allow-Headers", "accept", "authorization", "content-type"},
				AllowCredentials: true,
				MaxAge:           12 * 3600,
			}))

			r.GET("/v1/ws", handler.ChatHandler)
			r.POST("/v1/send-message", handler.SendMessageHandler)

			// Start the server
			addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
			fmt.Printf("Starting server for chat service on %s...\n", addr)

			if err := r.Run(addr); err != nil {
				fmt.Printf("Error starting server: %s\n", err)
				return
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return nil
	}

	return r
}
