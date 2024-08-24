package routes

import (
	"middleware/controllers"
	"middleware/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	// Protected routes
	r.Use(middlewares.AuthMiddleware)
	r.HandleFunc("/api/v1/packages", controllers.GetPackages).Methods("GET")
	r.HandleFunc("/api/v1/packages/{id}", controllers.GetPackageByID).Methods("GET")
	r.HandleFunc("/api/v1/packages", controllers.CreatePackage).Methods("POST")
	r.HandleFunc("/api/v1/packages/{id}", controllers.UpdatePackage).Methods("PUT")
	r.HandleFunc("/api/v1/packages/{id}", controllers.DeletePackage).Methods("DELETE")

	return r
}
