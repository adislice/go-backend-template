package user_repository

import (
	"context"

	user_model "github.com/adislice/go-project-structure/internal/modules/user/model"
)

func (r userRepository) GetAllUser(ctx context.Context) ([]user_model.User, error) {
	var users []user_model.User
	result := r.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []user_model.User{}, nil
	}

	return users, nil
}

func (r userRepository) GetUserById(ctx context.Context, id string) (user_model.User, error) {
	var user user_model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

	return user, err
}
