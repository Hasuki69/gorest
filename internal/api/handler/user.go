package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorest/internal/model/response"
	"gorm.io/gorm"
	"net/http"
)

func (h *Handler) GetAllUsers(ctx echo.Context) error {
	c := ctx.Request().Context()

	users, err := h.userService.GetAllUsers(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, response.Json{
				Message: "Not Found",
			})
		} else {
			return ctx.JSON(http.StatusInternalServerError, response.Json{
				Message: "Internal Server Error",
			})
		}
	}

	return ctx.JSON(http.StatusOK, response.Json{
		Message: "Success",
		Data:    &users,
	})
}
