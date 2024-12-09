package handlers

import (
	"chat-cache/cache"

	database "chat-database/db"
)

type Handler struct {
	Database *database.Database
	Cache    *cache.Cache
}
