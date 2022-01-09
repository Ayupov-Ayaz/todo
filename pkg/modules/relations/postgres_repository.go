package relations

import (
	"context"
	"database/sql"
	"errors"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

var (
	ErrListNotFound = _errors.NotFound("list not found")
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) GetListUserByListId(_ context.Context, listID int) (models.ListUser, error) {
	const get = `SELECT list_id, user_id 
				FROM users_lists 
				WHERE list_id = $1;`

	var listUser models.ListUser

	if err := p.db.Get(&listUser, get, listID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrListNotFound
		}

		return models.ListUser{}, err
	}

	return listUser, nil
}
