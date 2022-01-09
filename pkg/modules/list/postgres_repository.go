package list

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrListNotFound = _errors.NotFound("list not found or list does not belong to you")
)

const (
	create       = `INSERT INTO todo_list (title, description) VALUES ($1, $2) RETURNING id;`
	linkListUser = `INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2);`
	getAll       = `SELECT tl.id, tl.title 
					FROM todo_list tl 
					INNER JOIN users_lists ul on tl.id = ul.list_id 
					WHERE ul.user_id = $1;`
	getList = `SELECT tl.id, tl.title, tl.description 
					FROM todo_list tl
					INNER JOIN users_lists ul on tl.id = ul.list_id
					WHERE ul.user_id = $1 AND ul.list_id = $2;`
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
		_ = tx.Rollback()
		return 0, fmt.Errorf("insert list failed: %w", err)
	}

	if _, err := tx.Exec(linkListUser, userID, id); err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("insert link listID => userID failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("close transaction failed: %w", err)
	}

	return id, nil
}

func (r *PostgresRepository) GetAll(_ context.Context, userID int) ([]models.TodoList, error) {
	var lists []models.TodoList

	err := r.db.Select(&lists, getAll, userID)

	return lists, err
}

func (r *PostgresRepository) Get(_ context.Context, userID int, listID int) (models.TodoList, error) {
	var list models.TodoList

	err := r.db.Get(&list, getList, userID, listID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = ErrListNotFound
	}

	return list, err
}
