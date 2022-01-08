package list

import (
	"context"
	"fmt"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	create       = `INSERT INTO todo_list (title, description) VALUES ($1, $2) RETURNING id;`
	linkListUser = `INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2);`
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(_ context.Context, userID int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("create transaction failed: %w", err)
	}

	id := 0

	if err := tx.QueryRow(create, list.Title, list.Description).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert list failed: %w", err)
	}

	if err := tx.QueryRow(linkListUser, userID, id).Err(); err != nil {
		return 0, fmt.Errorf("insert link listID => userID failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("close transaction failed: %w", err)
	}

	return id, nil
}
