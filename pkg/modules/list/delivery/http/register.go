package http

import (
	"github.com/ayupov-ayaz/todo/pkg/modules/list"
	"github.com/gofiber/fiber/v2"
)

func RegisterHttpEndpoints(router fiber.Router, uc list.UseCase) {
	h := NewHandler(uc)

	group := router.Group("/lists")
	group.Post("/", h.Create)
	group.Get("/", h.GetLists)
	group.Get("/:id", h.Get)
	group.Patch("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
