package migrate

import (
	"context"
	"database/sql"
)

type Migration struct {
	DB *sql.DB
}

func (m *Migration) Up(ctx context.Context) error {
	_, err := m.DB.ExecContext(ctx, `ALTER TABLE authorizes DROP COLUMN jwt_token;`)
	return err
}

func (m *Migration) Down(ctx context.Context) error {
    _, err := m.DB.ExecContext(ctx, `ALTER TABLE authorizes ADD COLUMN jwt_token VARCHAR(255);`)
    return err
}