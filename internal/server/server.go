package server

import (
	"github.com/ayupov-ayaz/todo/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewServer(parser TokenParser) *fiber.App {
	cfg := http.Cfg{
		ReadTimeout:  viper.GetDuration("server.timeouts.read"),
		WriteTimeout: viper.GetDuration("server.timeouts.write"),
	}

	f := http.NewServer(cfg)
	f.Use(
		recovery(),
		auth(parser),
	)

	return f
}
