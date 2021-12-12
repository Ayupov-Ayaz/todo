package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	defaultPort         = 8080
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

type Cfg struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func DefaultCfg() Cfg {
	return Cfg{
		Port:         defaultPort,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
}

func NewServer(cfg Cfg) *fiber.App {
	return fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
}
