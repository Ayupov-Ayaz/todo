package app

import (
	"github.com/ayupov-ayaz/todo/pkg/modules/list/delivery/http"
	"github.com/ayupov-ayaz/todo/pkg/modules/list/repository"
	"github.com/ayupov-ayaz/todo/pkg/modules/list/usecase"
	"github.com/ayupov-ayaz/todo/pkg/services/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func listHandler(s *fiber.App, db *sqlx.DB, val validator.Validator) {
	repo := repository.NewPostgresRepository(db)
	srv := usecase.NewUseCase(repo, val)
	handler := http.NewHandler(srv)
	handler.RunHandler(s)

}
