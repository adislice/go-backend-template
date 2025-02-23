package user_repository

import (
	"context"

	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]user_model.User, error)
	GetUserById(ctx context.Context, id string) (user_model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
