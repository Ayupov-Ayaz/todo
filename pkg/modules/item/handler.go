package item

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func (h *Handler) RunHandler(router fiber.Router) {
	group := router.Group("/lists")

	group.Post("/", h.Create)
	group.Get("/", h.GetList)
	group.Get("/:id", h.Get)
	group.Patch("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}

func (h Handler) Create(ctx *fiber.Ctx) error {

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
