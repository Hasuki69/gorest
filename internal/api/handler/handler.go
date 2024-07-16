package handler

import (
	"gorest/internal/repository"
	"gorest/internal/service"
	"gorm.io/gorm"
)

type (
	Handler struct {
		userService service.UserService
	}
)

func NewHandler(db *gorm.DB) *Handler {

	// Create new user repository and service
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	return &Handler{
		userService: userService,
	}
}
