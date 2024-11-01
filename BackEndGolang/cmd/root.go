package cmd

import (
	"chat-backend/config"
	database "chat-backend/internal/db"
	"chat-backend/internal/handlers/v1"
	"chat-backend/internal/services"
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Execute() {
	configPath := filepath.Join("config", "config.yaml")
	config, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := database.GetClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	chatService := services.NewChatService()
	defer client.Close()
	rootCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the server",
		Run: func(cmd *cobra.Command, arg []string) {
			r := gin.Default()
			r.Use(cors.New(cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With", "X-Custom-Header", "Access-Control-Allow-Headers", "accept", "authorization", "content-type"},
				AllowCredentials: true,
				MaxAge:           12 * 3600,
			}))

			r.GET("/ws", func(ctx *gin.Context) {
				handlers.ChatHandler(ctx, chatService, client, config.Server.SecretKey)
			})
			r.POST("/v1/register", func(ctx *gin.Context) {
				handlers.RegisterHandler(ctx, client, config.Server.SecretKey)
			})
			r.POST("/v1/login", func(ctx *gin.Context) {
				handlers.LoginHandler(ctx, client, config.Server.SecretKey)
			})
			r.POST("/v1/send-message", func(ctx *gin.Context) {
				handlers.SendMessageHandler(ctx, client, chatService, config.Server.SecretKey)
			})
			r.GET("/v1/search", func(ctx *gin.Context) {
				handlers.FindUsersByNameHandler(ctx, client)
			})
			r.GET("/v1/rooms", func(ctx *gin.Context) {
				handlers.GetRoomsByUserHandler(ctx, client, config.Server.SecretKey)
			})
			r.GET("/v1/user-info", func(ctx *gin.Context) {
				handlers.GetUserInfoHandler(ctx, client, config.Server.SecretKey)
			})

			addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

			fmt.Printf("Starting server on %s...\n", addr)

			if err := r.Run(addr); err != nil {
				fmt.Printf("Error starting server: %s\n", err)
				return
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
