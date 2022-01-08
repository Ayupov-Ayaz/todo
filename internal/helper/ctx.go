package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrCastUserIdToIntFailed = errors.New("cast user id to int failed")
	ErrInvalidUserId         = errors.New("invalid user id")
)

const (
	userIdCtx = "userID"
)

func SetUserID(ctx *fiber.Ctx, userID int) {
	ctx.Locals(userIdCtx, userID)
}

func GetUserID(ctx *fiber.Ctx) (int, error) {
	id, ok := ctx.Locals(userIdCtx).(int)
	if !ok {
		return 0, ErrCastUserIdToIntFailed
	}

	if id <= 0 {
		return 0, ErrInvalidUserId
	}

	return id, nil
}
