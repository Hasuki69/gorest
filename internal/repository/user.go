package repository

import (
	"gorest/internal/model/domain"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetAllUsers() ([]*domain.User, error)
		// Add other methods here
	}

	userRepository struct {
		db *gorm.DB
	}
)

func (u userRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
