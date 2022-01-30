package http

import (
	"encoding/json"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	"github.com/ayupov-ayaz/todo/internal/models"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uc     item.UseCase
	logger *zap.Logger
}

func NewHandler(uc item.UseCase) *Handler {
	return &Handler{
		uc:     uc,
		logger: zap.L().Named("item_handler"),
	}
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

	id, err := h.uc.Create(ctx.UserContext(), userID, listID, item)
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

type getAllItemsResponse struct {
	Items []models.Item `json:"items"`
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

	items, err := h.uc.GetAll(ctx.UserContext(), userID, listID)
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

	item, err := h.uc.Get(ctx.UserContext(), userID, itemID)
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

func (h *Handler) Delete(ctx *fiber.Ctx) error {
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

	if err := h.uc.Delete(ctx.UserContext(), userID, itemID); err != nil {
		return err
	}

	http.SendOk(ctx, 200)

	return nil
}
