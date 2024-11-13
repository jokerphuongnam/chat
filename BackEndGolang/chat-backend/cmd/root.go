package cmd

import (
	"chat-backend/config"
	"chat-backend/internal/cache"
	database "chat-backend/internal/db"
	"chat-backend/internal/handlers/v1"
	"chat-backend/internal/logs"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func Execute(config config.AppConfig) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{config.ApiGateway.Addr})
	databaseHandler, err := database.GetClient(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	handler := &handlers.Handler{
		Database: databaseHandler,
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

			// Start the token cleanup scheduler with the database's GetAllUsersId function
			handler.Cache.StartTokenCleanupScheduler(func() ([]uuid.UUID, error) {
				return handler.Database.GetAllUsersId()
			})
			
			r.POST("/v1/send-message", proxyHandler(fmt.Sprintf("%v", config.ApiGateway.Addr)))

			r.POST("/v1/register", handler.RegisterHandler)
			r.POST("/v1/login", handler.LoginHandler)
			r.GET("/v1/search", handler.FindUsersByNameHandler)
			r.GET("/v1/rooms", handler.GetRoomsByUserHandler)
			r.GET("/v1/user-info", handler.GetUserInfoHandler)
			r.GET("/v1/rooms/:room_id/messages", handler.GetMessagesHandler)
			r.GET("/v1/room/room_id/info", handler.GetRoomInfoHandler)
			r.GET("/v1/room/get_room_id/:user_id", handler.GetRoomIDFromUserID)

			// Start the server
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
		return nil
	}

	return r
}

func proxyHandler(target string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logs.Log.Infof("Proxy Handler: %v", target)
		url, err := url.Parse(target)
		if err != nil {
			logs.Log.Infof("Invalid proxy target: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Invalid proxy target"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
