package app

import (
	"github.com/ayupov-ayaz/todo/pkg/modules/item/delivery/http"
	"github.com/ayupov-ayaz/todo/pkg/modules/item/repository"
	"github.com/ayupov-ayaz/todo/pkg/modules/item/usecase"
	"github.com/ayupov-ayaz/todo/pkg/modules/relations"
	"github.com/ayupov-ayaz/todo/pkg/services/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func itemHandler(s *fiber.App, db *sqlx.DB, val validator.Validator) {
	repo := repository.NewPostgresRepository(db)
	relationRepo := relations.NewPostgresRepository(db)
	srv := usecase.NewUseCase(repo, relationRepo, val)
	handler := http.NewHandler(srv)
	handler.RunHandler(s)
}
