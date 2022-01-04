package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Cfg struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(cfg Cfg) *fiber.App {
	return fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
}
