package item

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
	ErrItemNotFound = _errors.NotFound("item not found")
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) Create(_ context.Context, listID int, item models.Item) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	const create = `INSERT INTO todo_item (title, description) VALUES ($1, $2) RETURNING id;`

	if err := tx.QueryRow(create, item.Title, item.Description).Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("create item failed: %w", err)
	}

	const linkListItem = `INSERT INTO list_items (list_id, item_id) VALUES ($1, $2);`
	if _, err := tx.Exec(linkListItem, listID, id); err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("create relations item -> list failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("close transaction failed: %w", err)
	}

	return id, err
}

func (p *PostgresRepository) GetAll(_ context.Context, listID int) ([]models.Item, error) {
	const getAll = `SELECT ti.id, ti.title, ti.done
					FROM todo_item ti 
					INNER JOIN list_items li on ti.id = li.item_id
					WHERE li.list_id = $1;`

	var items []models.Item

	if err := p.db.Select(&items, getAll, listID); err != nil {
		return nil, err
	}

	return items, nil
}

func (p *PostgresRepository) Get(_ context.Context, itemID int) (models.Item, error) {
	const get = `SELECT id, title, description, done FROM todo_item WHERE id = $1;`

	var item models.Item

	if err := p.db.Get(&item, get, itemID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Item{}, ErrItemNotFound
		}
	}

	return item, nil
}

func (p *PostgresRepository) Delete(ctx context.Context, userID, itemID int) error {
	const query = `DELETE FROM todo_item ti 
						USING list_items li, users_lists ul 
					WHERE ti.id = li.item_id 
						AND li.list_id = ul.list_id
						AND ul.user_id = $1 AND ti.id = $2;`

	if _, err := p.db.Exec(query, userID, itemID); err != nil {
		return err
	}

	return nil
}
