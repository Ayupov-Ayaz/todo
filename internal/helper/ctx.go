package helper

import "github.com/gofiber/fiber/v2"

const (
	userIdCtx = "userID"
)

func SetUserID(ctx *fiber.Ctx, userID int) {
	ctx.Locals(userIdCtx, userID)
}

func GetUserID(ctx *fiber.Ctx) int {
	return ctx.Locals(userIdCtx).(int)
}
