package user_service

import (
	"context"
	"fmt"

	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
	pkg_error "github.com/adislice/go-project-structure/pkg/error"
	"github.com/adislice/go-project-structure/pkg/logger"
)

func (s userService) GetAllUser(ctx context.Context) ([]user_model.UserResponse, error) {
	users, err := s.userRepository.GetAllUser(ctx)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to get all users: %v", err))
		return nil, pkg_error.InternalServerError(err.Error())
	}

	var userResponses []user_model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user_model.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return userResponses, nil
}
