package http

import (
	"encoding/json"
	"errors"

	"github.com/dgrijalva/jwt-go"

	_errors "github.com/ayupov-ayaz/todo/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const (
	statusForbidden  = 403
	statusNotFound   = 404
	statusBadRequest = 400
)

func ParseError(err error) (message string, httpStatus int) {
	message = "internal error"
	httpStatus = 500

	var httpStatusErr _errors.HttpStatusError
	if errors.As(err, &httpStatusErr) {
		httpStatus = httpStatusErr.HttpStatus()
		message = httpStatusErr.Error()
		return
	}

	var jwtErr *jwt.ValidationError
	if errors.As(err, &jwtErr) {
		httpStatus = statusForbidden
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrForbidden) {
		httpStatus = statusForbidden
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrNotFound) {
		httpStatus = statusNotFound
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrBadRequest) {
		httpStatus = statusBadRequest
		message = err.Error()
		return
	}

	return
}

func ErrorHandler() func(ctx *fiber.Ctx, err error) error {
	logger := zap.L().Named("error_handler")

	return func(ctx *fiber.Ctx, err error) error {
		message, status := ParseError(err)

		raw, err := json.Marshal(&struct {
			Message string `json:"message"`
		}{
			Message: message,
		})

		if err != nil {
			logger.Error("marshaling error struct failed", zap.Error(err))
		}

		SendJson(ctx, raw, status)

		return nil
	}
}
