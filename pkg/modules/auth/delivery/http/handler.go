package http

import (
	"encoding/json"

	"github.com/ayupov-ayaz/todo/pkg/modules/auth"

	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/internal/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uc     auth.UseCse
	logger *zap.Logger
}

func NewHandler(uc auth.UseCse) *Handler {
	return &Handler{
		uc:     uc,
		logger: zap.L().Named("auth_handler"),
	}
}

func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		h.logger.Warn("unmarshal body failed",
			zap.ByteString("body", ctx.Body()),
			zap.Error(err))

		return auth.ErrInvalidRequest
	}

	id, err := h.uc.Create(user)
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

		return auth.ErrInvalidRequest
	}

	token, err := h.uc.SignIn(input.Username, input.Password)
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
