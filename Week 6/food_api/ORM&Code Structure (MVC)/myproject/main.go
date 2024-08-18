// main.go
package main

import (
	"myproject/config" // Adjust import path accordingly
	"myproject/routes" // Adjust import path accordingly

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Connect to the database
	config.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(e)

	// Start the server
	e.Start(":8080")
}
