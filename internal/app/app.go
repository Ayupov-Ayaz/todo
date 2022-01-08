package app

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/server"

	"github.com/ayupov-ayaz/todo/pkg/services/jwt"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"

	"github.com/joho/godotenv"

	"github.com/ayupov-ayaz/todo/pkg/repository"

	"github.com/spf13/viper"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/gofiber/fiber/v2"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth"
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

func authModel(
	s *fiber.App, db *sqlx.DB, jwtSrv jwt.Service, val validator.Validator,
	salt []byte, lifetime time.Duration,
) {
	repo := auth.NewPostgresRepository(db)
	srv := auth.NewService(repo, val, jwtSrv, salt, lifetime)
	handler := auth.NewHandler(srv)
	handler.RunHandler(s)
}

func itemHandler(s *fiber.App, db *sqlx.DB) {
	repo := item.NewPostgresRepository(db)
	srv := item.NewHandler(repo)
	handler := item.NewHandler(srv)
	handler.RunHandler(s)
}

func listHandler(s *fiber.App, db *sqlx.DB, val validator.Validator) {
	repo := list.NewPostgresRepository(db)
	srv := list.NewService(repo, val)
	handler := list.NewHandler(srv)
	handler.RunHandler(s)

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

	authCfg := auth.Config{
		Salt:       []byte(os.Getenv("PASS_SALT")),
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
		LifeTime:   viper.GetDuration("auth.lifetime"),
	}

	validate := validator.NewBasicValidator()

	if err := validate.Struct(authCfg); err != nil {
		return err
	}

	db, err := makePostgres()
	if err != nil {
		return err
	}

	jwtSrv := jwt.NewService(authCfg.SigningKey)
	s := server.NewServer(jwtSrv)

	authModel(s, db, jwtSrv, validate, authCfg.Salt, authCfg.LifeTime)
	listHandler(s, db, validate)
	itemHandler(s, db)

	if err := s.Listen(":" + strconv.Itoa(viper.GetInt("server.port"))); err != nil {
		return fmt.Errorf("occured while running http server: %w", err)
	}

	return nil
}
