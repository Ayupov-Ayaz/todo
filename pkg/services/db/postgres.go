package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSlMode  string
}

func MakePostgresDb(cfg PostgresConfig) (*sqlx.DB, error) {
	dbSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSlMode)

	db, err := sqlx.Open("postgres", dbSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
