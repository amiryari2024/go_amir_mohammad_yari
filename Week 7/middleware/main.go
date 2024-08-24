package main

import (
    "log"
    "net/http"
    "middleware/config"
    "middleware/models"
    "middleware/routes"
    "middleware/middlewares"
)

func main() {
    config.ConnectDatabase()
    models.Migrate(config.DB)

    r := routes.SetupRoutes()
    r.Use(middlewares.LoggingMiddleware)

    log.Println("Server starting on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
