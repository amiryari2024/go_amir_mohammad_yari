package main

import (
	"middleware/config"
	"middleware/routes"
)

func main() {
	config.InitDB()
	e := routes.InitRoutes()
	e.Start(":8000")
}
