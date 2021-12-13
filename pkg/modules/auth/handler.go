package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationService interface {
}

type Handler struct {
	srv AuthorizationService
}

func NewHandler(srv AuthorizationService) *Handler {
	return &Handler{srv: srv}
}

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
