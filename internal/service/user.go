package service

import (
	"context"
	"gorest/internal/model/domain"
	"gorest/internal/repository"
	"gorm.io/gorm"
	"time"
)

type (
	UserService interface {
		GetAllUsers(ctx context.Context) ([]*domain.User, error)
		// Add other methods here
	}

	userService struct {
		userRepository repository.UserRepository
	}
)

func (u userService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	_, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}
