package handlers

import (
	"chat-service/internal/cache"
	database "chat-service/internal/db"
	"chat-service/internal/services"
)

type Handler struct {
	ChatService *services.ChatService
	Database    *database.Database
	Cache       *cache.Cache
}
