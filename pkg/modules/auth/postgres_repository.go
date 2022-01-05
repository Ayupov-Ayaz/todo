package auth

import (
	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	create = "INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id"
	exist  = "SELECT EXISTS(SELECT id FROM users WHERE username = $1);"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(user models.User) (int, error) {
	id := 0

	if err := r.db.QueryRow(create, user.Name, user.Username, user.Password).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepository) Exist(username string) (bool, error) {
	ok := false
	if err := r.db.QueryRow(exist, username).Scan(&ok); err != nil {
		return false, err
	}

	return ok, nil
}
