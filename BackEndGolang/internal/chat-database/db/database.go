package database

import (
	"context"
	"database/sql"
	"fmt"

	"chat-database/ent"

	config "chat-config/config"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Client *ent.Client
	db     *sql.DB
}

func GetClient(cfg config.AppConfig) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to mysql: %w", err)
	}

	// Initialize Ent client
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to mysql: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}
	return &Database{
		Client: client,
		db:     db,
	}, nil
}
