package handlers

import (
	"chat-backend/internal/cache"

	database "../chat-database/db"
)

type Handler struct {
	Database *database.Database
	Cache    *cache.Cache
}
