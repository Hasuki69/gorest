package api

import (
	"gorest/internal/api/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handlers *handler.Handler) *echo.Echo {
	// Create a new echo instance
	e := echo.New()

	// Configure the middleware
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the initial routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Define group for the API
	api := e.Group("/api")
	api.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, API!")
	})

	// Define other routes
	user := api.Group("/user")
	user.GET("", handlers.GetAllUsers)

	return e
}
