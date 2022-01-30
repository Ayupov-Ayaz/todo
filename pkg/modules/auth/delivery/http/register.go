package http

import (
	"github.com/ayupov-ayaz/todo/pkg/modules/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterHttpEndpoints(router fiber.Router, uc auth.UseCse) {
	group := router.Group("/auth")

	h := NewHandler(uc)

	group.Post("/sign-up", h.SignUp)
	group.Post("/sign-in", h.SignIn)
}
