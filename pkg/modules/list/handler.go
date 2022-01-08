package list

import (
	"context"
	"encoding/json"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	_errors "github.com/ayupov-ayaz/todo/errors"
	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	ErrInvalidRequest = _errors.BadRequest("invalid request")
)

type TodoListService interface {
	Create(ctx context.Context, userID int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userID int) ([]models.TodoList, error)
}

type Handler struct {
	srv    TodoListService
	logger *zap.Logger
}

func NewHandler(srv TodoListService) *Handler {
	return &Handler{
		srv:    srv,
		logger: zap.L().Named("list_handler"),
	}
}

func (h *Handler) RunHandler(router fiber.Router) {
	group := router.Group("/lists")

	group.Post("/", h.Create)
	group.Get("/", h.GetLists)
	group.Get("/:id", h.Get)
	group.Patch("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}

func (h *Handler) Create(ctx *fiber.Ctx) error {
	userID, err := helper.GetUserID(ctx)
	if err != nil {
		h.logger.Error("get user id from ctx failed", zap.Error(err))
		return err
	}

	var list models.TodoList
	if err := json.Unmarshal(ctx.Body(), &list); err != nil {
		h.logger.Warn("unmarshal body failed", zap.Error(err))
		return ErrInvalidRequest
	}

	id, err := h.srv.Create(ctx.UserContext(), userID, list)
	if err != nil {
		return err
	}

	raw, err := helper.MarshalingId(id)
	if err != nil {
		h.logger.Error("marshaling response failed", zap.Error(err))
	}

	http.SendJson(ctx, raw, 200)

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

	lists, err := h.srv.GetAll(ctx.UserContext(), userID)
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

func (h Handler) Get(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) Update(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) Delete(ctx *fiber.Ctx) error {

	return nil
}
