package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) RunHandler(router fiber.Router) {
	group := router.Group("/auth")

	group.Post("/sign-up", h.SignUp)
	group.Post("/sign-in", h.SignIn)
}

func (h Handler) SignUp(ctx *fiber.Ctx) error {
	fmt.Println("sign-up")

	return nil
}

func (h Handler) SignIn(ctx *fiber.Ctx) error {
	fmt.Println("sign-in")

	return nil
}
