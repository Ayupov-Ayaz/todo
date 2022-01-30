package _errors

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ParseValidatorErrors(valErrors validator.ValidationErrors) string {
	last := len(valErrors) - 1
	var message string
	for i, err := range valErrors {
		message += err.Field() + ":" + err.Tag()

		if i != last {
			message += ";"
		}
	}

	return message
}

func ParseError(err error) (message string, httpStatus int) {
	message = "internal error"
	httpStatus = 500

	var httpStatusErr HttpStatusError
	if errors.As(err, &httpStatusErr) {
		httpStatus = httpStatusErr.HttpStatus()
		message = httpStatusErr.Error()
		return
	}

	var jwtErr *jwt.ValidationError
	if errors.As(err, &jwtErr) {
		httpStatus = fiber.StatusForbidden
		message = err.Error()
		return
	}

	var valErrors validator.ValidationErrors
	if errors.As(err, &valErrors) {
		httpStatus = fiber.StatusBadRequest
		message = ParseValidatorErrors(valErrors)
		return
	}

	if errors.Is(err, fiber.ErrForbidden) {
		httpStatus = fiber.StatusForbidden
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrNotFound) {
		httpStatus = fiber.StatusNotFound
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrBadRequest) {
		httpStatus = fiber.StatusBadRequest
		message = err.Error()
		return
	}

	if errors.Is(err, fiber.ErrMethodNotAllowed) {
		httpStatus = fiber.StatusMethodNotAllowed
		message = err.Error()
		return
	}

	return
}
