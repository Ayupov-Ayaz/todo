package app

import (
	"fmt"
	"strconv"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/gofiber/fiber/v2"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"
)

func authModel(s *fiber.App) {
	repo := auth.NewRepository()
	srv := auth.NewHandler(repo)
	handler := auth.NewHandler(srv)
	handler.RunHandler(s)
}

func itemHandler(s *fiber.App) {
	repo := item.NewRepository()
	srv := item.NewHandler(repo)
	handler := item.NewHandler(srv)
	handler.RunHandler(s)
}

func listHandler(s *fiber.App) {
	repo := list.NewRepository()
	srv := list.NewHandler(repo)
	handler := list.NewHandler(srv)
	handler.RunHandler(s)

}

func Run() error {
	cfg := http.DefaultCfg()
	s := http.NewServer(cfg)

	authModel(s)
	listHandler(s)
	itemHandler(s)

	if err := s.Listen(":" + strconv.Itoa(cfg.Port)); err != nil {
		return fmt.Errorf("occured while running http server: %w", err)
	}

	return nil
}
