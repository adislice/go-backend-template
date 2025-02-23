package utils

import (
	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
	"github.com/gofiber/fiber/v2"
)

func GetAuthUser(ctx *fiber.Ctx) user_model.User {
	user, ok := ctx.Locals("user").(user_model.User)
	if !ok {
		return user_model.User{}
	}

	return user
}
