package handlers

import (
	"chat-backend/internal/cache"
	database "chat-backend/internal/db"
	"chat-backend/internal/services"
)

type Handler struct {
	ChatService *services.ChatService
	Database    *database.Database
	Cache       *cache.Cache
}
