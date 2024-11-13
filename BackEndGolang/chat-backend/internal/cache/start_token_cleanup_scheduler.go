package cache

import (
	"log"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

func (cache *Cache) StartTokenCleanupScheduler(getAllUsersId func() ([]uuid.UUID, error)) {
	c := cron.New()
	c.AddFunc("@daily", func() {
		users, err := getAllUsersId()
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return
		}
		for _, userID := range users {
			err := cache.CleanUpExpiredTokens(userID.String())
			if err != nil {
				log.Printf("Error cleaning up expired tokens for user %s: %v", userID, err)
			} else {
				log.Printf("Successfully cleaned up expired tokens for user %s", userID)
			}
		}
	})

	c.Start()
}
