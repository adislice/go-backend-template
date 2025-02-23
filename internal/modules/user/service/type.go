package user_service

import (
	"context"

	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
	user_repository "github.com/adislice/go-project-structure/internal/modules/user/repository"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]user_model.UserResponse, error)
}

type userService struct {
	userRepository user_repository.UserRepository
}

func NewUserService(userRepository user_repository.UserRepository) UserService {
	return userService{
		userRepository: userRepository,
	}
}
