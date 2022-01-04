package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Cfg struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func SendJson(ctx *fiber.Ctx, raw []byte, httpCode int) {
	ctx.Status(httpCode)
	ctx.Response().Header.SetContentType("application/json")
	ctx.Response().SetBodyRaw(raw)
}

func NewServer(cfg Cfg) *fiber.App {
	return fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		ErrorHandler: ErrorHandler(),
	})
}
