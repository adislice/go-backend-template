package user_route

import (
	"github.com/adislice/go-project-structure/internal/middleware"
	user_handler "github.com/adislice/go-project-structure/internal/modules/user/http/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, authMiddleware middleware.AuthMiddleware, userHandler user_handler.UserHandler) {
	userRoute := router.Group("/users")

	userRoute.Use(authMiddleware)
	userRoute.Get("/", userHandler.GetAllUser)
}
