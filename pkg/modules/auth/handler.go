package auth

import (
	"encoding/json"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/internal/models"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidRequest = _errors.BadRequest("invalid request")
)

type AuthorizationService interface {
	Create(user models.User) (int, error)
	SignIn(username, password string) (string, error)
}

type Handler struct {
	srv    AuthorizationService
	logger *zap.Logger
}

func NewHandler(srv AuthorizationService) *Handler {
	return &Handler{
		srv:    srv,
		logger: zap.L().Named("auth_handler"),
	}
}

func (h *Handler) RunHandler(router fiber.Router) {
	group := router.Group("/auth")

	group.Post("/sign-up", h.SignUp)
	group.Post("/sign-in", h.SignIn)
}

func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		h.logger.Warn("unmarshal body failed",
			zap.ByteString("body", ctx.Body()),
			zap.Error(err))

		return ErrInvalidRequest
	}

	id, err := h.srv.Create(user)
	if err != nil {
		return err
	}

	raw, err := helper.MarshalingId(id)
	if err != nil {
		h.logger.Error("marshaling response failed", zap.Error(err))
	}

	http.SendJson(ctx, raw, 201)

	return nil
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(ctx *fiber.Ctx) error {
	var input SignInInput

	if err := json.Unmarshal(ctx.Body(), &input); err != nil {
		h.logger.Error("unmarshal failed",
			zap.ByteString("body", ctx.Body()),
			zap.Error(err))

		return ErrInvalidRequest
	}

	token, err := h.srv.SignIn(input.Username, input.Password)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		Token: token,
	})

	if err != nil {
		h.logger.Error("marshaling response failed", zap.Error(err))
	}

	http.SendJson(ctx, raw, 200)

	return nil
}
