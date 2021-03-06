package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/jmoiron/sqlx"
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
	const (
		create       = `INSERT INTO todo_list (title, description) VALUES ($1, $2) RETURNING id;`
		linkListUser = `INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2);`
	)

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
	const getAll = `SELECT tl.id, tl.title 
					FROM todo_list tl 
					INNER JOIN users_lists ul on tl.id = ul.list_id 
					WHERE ul.user_id = $1;`

	var lists []models.TodoList

	err := r.db.Select(&lists, getAll, userID)

	return lists, err
}

func (r *PostgresRepository) Get(_ context.Context, userID int, listID int) (models.TodoList, error) {
	const getList = `SELECT tl.id, tl.title, tl.description 
					FROM todo_list tl
					INNER JOIN users_lists ul on tl.id = ul.list_id
					WHERE ul.user_id = $1 AND ul.list_id = $2;`

	var todoList models.TodoList

	err := r.db.Get(&todoList, getList, userID, listID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = list.ErrListNotFound
	}

	return todoList, err
}

func (r *PostgresRepository) Update(_ context.Context, userID, listID int, input list.UpdateTodoList) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, "title=$"+strconv.Itoa(argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, "description=$"+strconv.Itoa(argID))
		args = append(args, *input.Description)
		argID++
	}

	args = append(args, userID, listID)

	//title=$1
	//description=$1
	//title=$1, descriptions=$2
	setQuery := strings.Join(setValues, ", ")

	const update = `UPDATE todo_list tl SET %s 
				FROM users_lists ul 
				WHERE tl.id = ul.list_id 
				AND ul.user_id = $%d
				AND ul.list_id = $%d;`

	query := fmt.Sprintf(update, setQuery, argID, argID+1)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *PostgresRepository) Delete(_ context.Context, userID, listID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("create transaction failed: %w", err)
	}

	const deleteListItems = `DELETE FROM todo_item ti
							USING list_items li, users_lists ul
							WHERE li.item_id = ti.id
							AND ul.user_id = $1  
							AND li.list_id = $2;`

	_, err = tx.Exec(deleteListItems, userID, listID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete items failed: %w", err)
	}

	const deleteList = `DELETE FROM todo_list tl 
					USING users_lists ul 
					WHERE tl.id = ul.list_id 
					AND ul.user_id = $1 
					AND ul.list_id = $2;`

	_, err = tx.Exec(deleteList, userID, listID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete list failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("close transaction failed: %w", err)
	}

	return nil
}
