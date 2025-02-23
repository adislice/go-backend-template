package user_handler

import (
	user_service "github.com/adislice/go-project-structure/internal/modules/user/service"
	"github.com/adislice/go-project-structure/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAllUser(c *fiber.Ctx) error
}

type userHandler struct {
	validator   validation.Validator
	userService user_service.UserService
}

func NewUserHandler(validator validation.Validator, userService user_service.UserService) UserHandler {
	return userHandler{
		validator:   validator,
		userService: userService,
	}
}
