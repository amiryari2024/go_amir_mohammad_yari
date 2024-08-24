package controllers

import (
	"myproject/config"
	"myproject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPackages(c echo.Context) error {
	var packages []models.Package
	config.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
	id := c.Param("id")
	var pkg models.Package
	config.DB.First(&pkg, id)
	if pkg.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}
	return c.JSON(http.StatusOK, pkg)
}

func CreatePackage(c echo.Context) error {
	pkg := new(models.Package)
	if err := c.Bind(pkg); err != nil {
		return err
	}
	config.DB.Create(pkg)
	return c.JSON(http.StatusCreated, pkg)
}

func UpdatePackage(c echo.Context) error {
	id := c.Param("id")
	var pkg models.Package
	config.DB.First(&pkg, id)
	if pkg.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}
	if err := c.Bind(&pkg); err != nil {
		return err
	}
	config.DB.Save(&pkg)
	return c.JSON(http.StatusOK, pkg)
}

func DeletePackage(c echo.Context) error {
	id := c.Param("id")
	var pkg models.Package
	config.DB.First(&pkg, id)
	if pkg.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}
	config.DB.Delete(&pkg)
	return c.NoContent(http.StatusNoContent)
}
