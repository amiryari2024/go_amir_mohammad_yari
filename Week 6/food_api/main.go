package main

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

type Food struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

var foods []Food = []Food{}

func getallfoods(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data": foods,
	})
}

func getfoodbyid(c echo.Context) error {
	foodid := c.Param("id")
	var food Food
	for _, f := range foods {
		if f.ID == foodid {
			food = f

		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": food,
	})

}

func createfood(c echo.Context) error {
	var food Food
	c.Bind(&food)
	foods = append(foods, food)
	return c.JSON(http.StatusCreated, echo.Map{
		"data": food,
	})
}

func updatefood(c echo.Context) error {
	var food Food
	c.Bind(&food)
	foodid := c.Param("id")
	for i := 0; i < len(foods); i++ {
		if foods[i].ID == foodid {
			foods[i].Name = food.Name
			foods[i].Description = food.Description
			foods[i].Price = food.Price

		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": food,
	})

}

func deletefood(c echo.Context) error {
	foodid := c.Param("id")
	foods = slices.DeleteFunc(foods, func(f Food) bool {
		return f.ID == foodid

	})
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
	})
}

func main() {
	e := echo.New()
	e.GET("/api/v1/foods", getallfoods)
	e.GET("/api/v1/foods/:id", getfoodbyid)
	e.POST("/api/v1/foods", createfood)
	e.PUT("/api/v1/foods/:id", updatefood)
	e.DELETE("/api/v1/foods/:id", deletefood)

	e.Logger.Fatal(e.Start(":1323"))
}
