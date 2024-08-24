package controllers

import (
	"encoding/json"
	"net/http"
	"your_project/models"
)

func GetPackages(w http.ResponseWriter, r *http.Request) {
	var packages []models.Package
	models.DB.Find(&packages)
	json.NewEncoder(w).Encode(packages)
}

func GetPackageByID(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func CreatePackage(w http.ResponseWriter, r *http.Request) {
	var pkg models.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	models.DB.Create(&pkg)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pkg)
}

func UpdatePackage(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func DeletePackage(w http.ResponseWriter, r *http.Request) {
	// Implementation
}
