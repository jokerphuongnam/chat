package handlers

import (
	"chat-cache/cache"
	database "chat-database/db"
	"chat-service/internal/services"
)

type Handler struct {
	ChatService *services.ChatService
	Database    *database.Database
	Cache       *cache.Cache
}
