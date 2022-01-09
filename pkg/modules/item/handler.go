package item

import (
	"context"
	"encoding/json"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	"github.com/ayupov-ayaz/todo/internal/models"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type TodoItemService interface {
	Create(ctx context.Context, userID, listID int, item models.Item) (int, error)
	GetAll(ctx context.Context, userID, listID int) ([]models.Item, error)
	Get(ctx context.Context, userID, itemID int) (models.Item, error)
}

type Handler struct {
	srv    TodoItemService
	logger *zap.Logger
}

func NewHandler(srv TodoItemService) *Handler {
	return &Handler{
		srv:    srv,
		logger: zap.L().Named("item_handler"),
	}
}

func (h *Handler) RunHandler(router fiber.Router) {
	groupList := router.Group("/:listID/item")
	groupList.Post("/", h.Create)
	groupList.Get("/", h.GetAll)

	groupItem := router.Group("item/:itemID")
	groupItem.Get("", h.Get)
	groupItem.Patch("", h.Update)
	groupItem.Delete("", h.Delete)
}

func getListID(ctx *fiber.Ctx) (int, error) {
	return ctx.ParamsInt("listID")
}

func getItemID(ctx *fiber.Ctx) (int, error) {
	return ctx.ParamsInt("itemID")
}

func (h *Handler) Create(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	listID, err := getListID(ctx)
	if err != nil {
		h.logger.Warn("param list id doesn't send", zap.Error(err))
		return err
	}

	var item models.Item
	if err := json.Unmarshal(ctx.Body(), &item); err != nil {
		h.logger.Warn("unmarshal body failed", zap.Error(err))
		return err
	}

	id, err := h.srv.Create(ctx.UserContext(), userID, listID, item)
	if err != nil {
		return err
	}

	raw, err := helper.MarshalingId(id)
	if err != nil {
		h.logger.Error("marshaling id failed", zap.Error(err))
		return err
	}

	http.SendJson(ctx, raw, 201)

	return nil
}

func (h *Handler) GetAll(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	listID, err := getListID(ctx)
	if err != nil {
		h.logger.Warn("param list id doesn't send", zap.Error(err))
		return err
	}

	items, err := h.srv.GetAll(ctx.UserContext(), userID, listID)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(getAllItemsResponse{Items: items})
	if err != nil {
		h.logger.Error("marshaling get all items response", zap.Error(err))
		return err
	}

	http.SendJson(ctx, raw, 200)

	return nil
}

func (h *Handler) Get(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	itemID, err := getItemID(ctx)
	if err != nil {
		h.logger.Warn("param item id doesn't send", zap.Error(err))
		return err
	}

	item, err := h.srv.Get(ctx.UserContext(), userID, itemID)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(item)
	if err != nil {
		h.logger.Error("marshaling item failed", zap.Error(err))
		return err
	}

	http.SendJson(ctx, raw, 200)

	return nil
}

func (h Handler) Update(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) Delete(ctx *fiber.Ctx) error {

	return nil
}
