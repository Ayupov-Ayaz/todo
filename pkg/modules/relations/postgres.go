package relations

import (
	"context"
	"database/sql"
	"errors"

	_errors "github.com/ayupov-ayaz/todo/internal/errors"

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

func (p *PostgresRepository) getListUser(query string, args ...interface{}) (models.ListUserRelation, error) {
	var listUser models.ListUserRelation

	if err := p.db.Get(&listUser, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrListNotFound
		}

		return models.ListUserRelation{}, err
	}

	return listUser, nil

}

func (p *PostgresRepository) GetListUserRelationByListId(_ context.Context, listID int) (models.ListUserRelation, error) {
	const query = `SELECT list_id, user_id 
				FROM users_lists 
				WHERE list_id = $1;`

	return p.getListUser(query, listID)
}

func (p *PostgresRepository) GetListUserRelationByItemId(_ context.Context, itemID int) (models.ListUserRelation, error) {
	const query = `SELECT ul.list_id, ul.user_id FROM users_lists ul
					INNER JOIN list_items li ON li.list_id = ul.list_id
					WHERE li.item_id = $1;`

	return p.getListUser(query, itemID)
}
