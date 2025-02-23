package routes

import (
	"github.com/adislice/go-project-structure/internal/middleware"
	user_handler "github.com/adislice/go-project-structure/internal/modules/user/http/handler"
	user_route "github.com/adislice/go-project-structure/internal/modules/user/http/route"
	user_repository "github.com/adislice/go-project-structure/internal/modules/user/repository"
	user_service "github.com/adislice/go-project-structure/internal/modules/user/service"
	"github.com/adislice/go-project-structure/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApiRoutes(router fiber.Router, db *gorm.DB, validator validation.Validator) {
	// Inisiasi auth middleware
	authMiddleware := middleware.NewAuthMiddleware(db)

	// Inisiasi repository
	userRepository := user_repository.NewUserRepository(db)

	// Inisiasi service
	userService := user_service.NewUserService(userRepository)

	// Inisasi handler
	userHandler := user_handler.NewUserHandler(validator, userService)

	apiRoute := router.Group("/api")

	// Setup module's routes
	user_route.SetupRoutes(apiRoute, authMiddleware, userHandler)
}
