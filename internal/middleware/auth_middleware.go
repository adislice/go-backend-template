package middleware

import (
	"github.com/adislice/go-project-structure/config"
	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
	pkgerr "github.com/adislice/go-project-structure/pkg/error"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthMiddleware fiber.Handler

// AuthMiddleware authenticates the JWT token
func NewAuthMiddleware(db *gorm.DB) AuthMiddleware {
	return jwtware.New(jwtware.Config{
		ContextKey:     "token",
		SigningKey:     jwtware.SigningKey{Key: []byte(config.AppConfig.JWTSecretKey)},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess(db),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(pkgerr.BadRequest("Missing or malformed JWT"))
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgerr.Unauthorized("Unauthorized: Invalid token"))
	}
}

func jwtSuccess(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var userModel user_model.User

		user := c.Locals("token").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		idUser := claims["user_id"].(string)

		if err := db.Model(&user_model.User{}).Where("id = ?", idUser).First(&userModel).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(pkgerr.Unauthorized("Unauthorized: Invalid token"))
		}

		c.Locals("user", userModel)
		return c.Next()
	}
}
