package http

import (
	"github.com/ayupov-ayaz/todo/pkg/modules/item"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router, uc item.UseCase) {
	h := NewHandler(uc)

	groupList := router.Group("/:listID/item")
	groupList.Post("/", h.Create)
	groupList.Get("/", h.GetAll)

	groupItem := router.Group("item/:itemID")
	groupItem.Get("", h.Get)
	groupItem.Patch("", h.Update)
	groupItem.Delete("", h.Delete)
}
