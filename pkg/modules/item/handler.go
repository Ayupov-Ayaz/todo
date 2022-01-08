package item

import (
	"fmt"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type TodoItemService interface {
}

type Handler struct {
	srv TodoItemService
}

func NewHandler(srv TodoItemService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RunHandler(router fiber.Router) {
	group := router.Group("/item")

	group.Post("/", h.Create)
	group.Get("/", h.GetList)
	group.Get("/:id", h.Get)
	group.Patch("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}

func (h *Handler) Create(ctx *fiber.Ctx) error {
	userID := helper.GetUserID(ctx)
	fmt.Println(userID)

	return nil
}

func (h Handler) Get(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) GetList(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) Update(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) Delete(ctx *fiber.Ctx) error {

	return nil
}
