package database

import (
	"chat-backend/config"
	"chat-backend/internal/ent"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetClient(cfg config.AppConfig) (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	// Connect to the database
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to mysql: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return client, nil
}
