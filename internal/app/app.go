package app

import (
	"fmt"
	"strconv"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/gofiber/fiber/v2"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"
)

func authModel(s *fiber.App) {
	authHandler := new(auth.Handler)
	authHandler.RunHandler(s)
}

func itemHandler(s *fiber.App) {
	itemHandler := new(item.Handler)
	itemHandler.RunHandler(s)
}

func listHandler(s *fiber.App) {
	listHandler := new(list.Handler)
	listHandler.RunHandler(s)
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
