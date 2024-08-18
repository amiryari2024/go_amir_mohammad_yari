// routes/routes.go
package routes

import (
	"myproject/controllers" // Adjust import path accordingly

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/packages", controllers.GetPackages)
	e.GET("/api/v1/packages/:id", controllers.GetPackage)
	e.POST("/api/v1/packages", controllers.CreatePackage)
	e.PUT("/api/v1/packages/:id", controllers.UpdatePackage)
	e.DELETE("/api/v1/packages/:id", controllers.DeletePackage)
}
