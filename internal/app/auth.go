package app

import (
	"time"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth/delivery/http"
	"github.com/ayupov-ayaz/todo/pkg/modules/auth/repositories"
	"github.com/ayupov-ayaz/todo/pkg/modules/auth/usecase"
	"github.com/ayupov-ayaz/todo/pkg/services/jwt"
	"github.com/ayupov-ayaz/todo/pkg/services/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func authModel(
	s *fiber.App, db *sqlx.DB, jwtSrv jwt.Service, val validator.Validator,
	salt []byte, lifetime time.Duration,
) {
	repo := repositories.NewPostgresRepository(db)
	uc := usecase.NewUseCase(repo, val, jwtSrv, salt, lifetime)
	http.RegisterHttpEndpoints(s, uc)
}
