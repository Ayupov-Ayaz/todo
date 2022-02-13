package app

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ayupov-ayaz/todo/pkg/services/db"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/server"

	"github.com/ayupov-ayaz/todo/pkg/services/jwt"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"

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
	viper.SetConfigName(".config")

	return viper.ReadInConfig()
}

func makePostgres() (*sqlx.DB, error) {
	cfg := db.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSlMode:  viper.GetString("db.sslmode"),
	}

	return db.MakePostgresDb(cfg)
}

func Run() error {
	if err := initConfig(); err != nil {
		return err
	}

	if err := initLogger(); err != nil {
		return err
	}

	authCfg := auth.InitConfig()

	validate := validator.NewBasicValidator()

	if err := validate.Struct(authCfg); err != nil {
		return err
	}

	db, err := makePostgres()
	if err != nil {
		return err
	}

	jwtSrv := jwt.NewUseCase(authCfg.SigningKey)
	s := server.NewServer(jwtSrv)

	authModel(s, db, jwtSrv, validate, authCfg.Salt, authCfg.LifeTime)
	listHandler(s, db, validate)
	itemHandler(s, db, validate)

	l := zap.L().Named("app")
	go func() {
		if err := s.Listen(":" + strconv.Itoa(viper.GetInt("server.port"))); err != nil {
			l.Error("start server failed", zap.Error(err))
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	_type := <-shutdown

	l.Info("graceful shutdown", zap.String("type", _type.String()))

	if err := s.Shutdown(); err != nil {
		l.Error("failed to shutdown fiber server", zap.Error(err))
	}

	if err := db.Close(); err != nil {
		l.Error("failed to shutdown database")
	}

	return nil
}
