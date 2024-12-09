package database

import (
	"chat-database/internal/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type RegisterResponse struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	JwtToken string    `json:"token"`
}

func (db *Database) RegisterHandler(username, password, name string) (*RegisterResponse, error) {
	// Start transaction
	tx, err := db.Client.Tx(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // Rollback on panic
			panic(p)          // Re-throw the panic after rollback
		} else if err != nil {
			_ = tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	// Validate the username and password format
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password %w", err)
	}

	// Create a new username/password and authorize
	up, err := db.Client.UsernamePassword.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create username/password: %w", err)
	}

	// Create a new authorize
	au, err := db.Client.Authorize.Create().
		SetToken(up.ID.String()).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to create authorize: %w", err)
	}

	// Create a new user
	user, err := db.Client.User.Create().
		SetName(name).
		SetIDAuthorize(au.ID).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: up.Username,
	}, nil
}
