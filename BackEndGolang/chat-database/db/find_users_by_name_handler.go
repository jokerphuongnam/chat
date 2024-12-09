package database

import (
	"chat-database/ent/user"
	"context"
	"fmt"
	"strings"
)

type FindUsersByNameResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func (db *Database) FindUsersByNameHandler(term string) ([]FindUsersByNameResponse, error) {
	words := strings.Fields(term)
	query := db.Client.User.Query()
	for _, word := range words {
		query = query.Where(user.NameContainsFold(word))
	}

	users, err := query.All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to find users by name: %w", err)
	}

	var results []FindUsersByNameResponse
	if len(users) > 0 {
		for _, user := range users {
			results = append(results, FindUsersByNameResponse{
				ID:        user.ID.String(),
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
			})
		}
	} else {
		results = []FindUsersByNameResponse{}
	}

	return results, nil
}
