// controllers/package_controller.go
package controllers

import (
	"myproject/config" // Adjust import path accordingly
	"myproject/models" // Adjust import path accordingly
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get all packages
func GetPackages(c echo.Context) error {
	var packages []models.Package
	config.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

// Get a package by ID
func GetPackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package
	if err := config.DB.First(&packageData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}
	return c.JSON(http.StatusOK, packageData)
}

// Add a new package
func CreatePackage(c echo.Context) error {
	var packageData models.Package
	if err := c.Bind(&packageData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Create(&packageData)
	return c.JSON(http.StatusCreated, packageData)
}

// Update package data
func UpdatePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package
	if err := config.DB.First(&packageData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}

	if err := c.Bind(&packageData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	config.DB.Save(&packageData)
	return c.JSON(http.StatusOK, packageData)
}

// Delete a package
func DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var packageData models.Package
	if err := config.DB.Delete(&packageData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Package deleted successfully"})
}
