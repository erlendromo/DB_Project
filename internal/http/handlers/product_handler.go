package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"net/http"
)

// GetAllProducts returns a slice of all products in the database
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := dependencies.Dependencies.ProductDeps.ProductDomain.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, products)
}
