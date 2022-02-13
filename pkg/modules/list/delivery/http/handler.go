package http

import (
	"encoding/json"

	"github.com/ayupov-ayaz/todo/pkg/modules/list"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	uc     list.UseCase
	logger *zap.Logger
}

func NewHandler(uc list.UseCase) *Handler {
	return &Handler{
		uc:     uc,
		logger: zap.L().Named("list_handler"),
	}
}

func (h *Handler) Create(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	var todoList models.TodoList
	if err := json.Unmarshal(ctx.Body(), &todoList); err != nil {
		h.logger.Warn("unmarshal body failed", zap.Error(err))
		return list.ErrInvalidRequest
	}

	id, err := h.uc.Create(ctx.UserContext(), userID, todoList)
	if err != nil {
		return err
	}

	http.SendJson(ctx, helper.MarshalingId(id), 200)

	return nil
}

type getAllListResponse struct {
	List []models.TodoList `json:"list"`
}

func (h *Handler) GetLists(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	lists, err := h.uc.GetAll(ctx.UserContext(), userID)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(getAllListResponse{List: lists})
	if err != nil {
		h.logger.Error("marshaling lists failed", zap.Error(err))
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

	listID, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Warn("param id doesn't send", zap.Error(err))
		return err
	}

	todoList, err := h.uc.Get(ctx.UserContext(), userID, listID)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(todoList)
	if err != nil {
		h.logger.Error("marshaling list failed", zap.Error(err))
		return err
	}

	http.SendJson(ctx, raw, 200)

	return nil
}

func (h *Handler) Update(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	listID, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Warn("param id doesn't send", zap.Error(err))
		return err
	}

	var input list.UpdateTodoList
	if err := json.Unmarshal(ctx.Body(), &input); err != nil {
		h.logger.Warn("unmarshal body failed", zap.Error(err))
		return list.ErrInvalidRequest
	}

	if err := h.uc.Update(ctx.UserContext(), userID, listID, input); err != nil {
		return err
	}

	http.SendOk(ctx, 200)

	return nil
}

func (h *Handler) Delete(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	listID, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Warn("param id doesn't send", zap.Error(err))
		return err
	}

	if err := h.uc.Delete(ctx.UserContext(), userID, listID); err != nil {
		return err
	}

	http.SendOk(ctx, 200)

	return nil
}
