package repository

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

func makeSqlxDB(cfg PostgresConfig) (*sqlx.DB, error) {
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

type PostgresDb struct {
	sqlx *sqlx.DB
}

func MakePostgresDb(cfg PostgresConfig) (*PostgresDb, error) {
	db, err := makeSqlxDB(cfg)
	if err != nil {
		return nil, err
	}

	return &PostgresDb{sqlx: db}, nil
}
