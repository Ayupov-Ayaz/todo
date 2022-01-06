package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"

	"github.com/joho/godotenv"

	"github.com/ayupov-ayaz/todo/pkg/repository"

	"github.com/spf13/viper"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/gofiber/fiber/v2"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"
)

const (
	envFilePath = "./configs/.env"
)

func initLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)

	return nil
}

func initConfig() error {
	if err := godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("error loading env variables: %w", err)
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func authModel(s *fiber.App, db *sqlx.DB) error {
	repo := auth.NewPostgresRepository(db)
	cfg := auth.Config{
		Salt:       []byte(os.Getenv("PASS_SALT")),
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
		Expired:    viper.GetDuration("auth.expired"),
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	srv := auth.NewService(repo, cfg)
	handler := auth.NewHandler(srv)
	handler.RunHandler(s)

	return nil
}

func itemHandler(s *fiber.App, db *sqlx.DB) {
	repo := item.NewPostgresRepository(db)
	srv := item.NewHandler(repo)
	handler := item.NewHandler(srv)
	handler.RunHandler(s)
}

func listHandler(s *fiber.App, db *sqlx.DB) {
	repo := list.NewPostgresRepository(db)
	srv := list.NewHandler(repo)
	handler := list.NewHandler(srv)
	handler.RunHandler(s)

}

func makeServer() *fiber.App {
	cfg := http.Cfg{
		ReadTimeout:  viper.GetDuration("server.timeouts.read"),
		WriteTimeout: viper.GetDuration("server.timeouts.write"),
	}

	return http.NewServer(cfg)
}

func makePostgres() (*sqlx.DB, error) {
	cfg := repository.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSlMode:  viper.GetString("db.sslmode"),
	}

	return repository.MakePostgresDb(cfg)
}

func Run() error {
	if err := initConfig(); err != nil {
		return err
	}

	if err := initLogger(); err != nil {
		return err
	}

	db, err := makePostgres()
	if err != nil {
		return err
	}

	s := makeServer()

	if err := authModel(s, db); err != nil {
		return err
	}

	listHandler(s, db)
	itemHandler(s, db)

	if err := s.Listen(":" + strconv.Itoa(viper.GetInt("server.port"))); err != nil {
		return fmt.Errorf("occured while running http server: %w", err)
	}

	return nil
}
