package server

import (
	"strings"

	_errors "github.com/ayupov-ayaz/todo/errors"
	"github.com/ayupov-ayaz/todo/internal/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	authorizationHeader = "Authorization"
)

type TokenParser interface {
	ParseToken(accessToken string) (userID int, err error)
}

func skip(ctx *fiber.Ctx) bool {
	uri := string(ctx.Request().URI().Path())

	return ctx.Route().Method == fiber.MethodGet || strings.HasPrefix(uri, "/auth")
}

func auth(parser TokenParser) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if !skip(ctx) {
			header := ctx.Get(authorizationHeader)
			if header == "" {
				return _errors.Forbidden("empty auth header")
			}

			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 {
				return _errors.Forbidden("invalid auth header")
			}

			userID, err := parser.ParseToken(headerParts[1])
			if err != nil {
				return err
			}

			helper.SetUserID(ctx, userID)
		}

		return ctx.Next()
	}
}

func recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
	})
}
