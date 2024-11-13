package handlers

import (
	"chat-backend/internal/cache"
	database "chat-backend/internal/db"
)

type Handler struct {
	Database *database.Database
	Cache    *cache.Cache
}
