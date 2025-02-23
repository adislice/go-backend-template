package user_handler

import (
	pkg_error "github.com/adislice/go-project-structure/pkg/error"
	pkg_success "github.com/adislice/go-project-structure/pkg/success"
	"github.com/gofiber/fiber/v2"
)

func (h userHandler) GetAllUser(c *fiber.Ctx) error {
	res, err := h.userService.GetAllUser(c.Context())
	if err != nil {
		return pkg_error.HandleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(pkg_success.SuccessWithData(res))
}
