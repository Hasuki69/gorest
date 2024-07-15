package main

import (
	"fmt"
	"gorest/internal/api"
	"gorest/internal/api/handler"
	"gorest/pkg/config"
)

func main() {
	// Load the config file
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize the database
	db, err := config.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize the handler
	handlers := handler.NewHandler(db)

	// Initialize echo and router
	echo := api.NewRouter(handlers)

	// Start the server on the specified port
	echo.Logger.Fatal(echo.Start(fmt.Sprintf(":%d", cfg.Server.Port)))

}
