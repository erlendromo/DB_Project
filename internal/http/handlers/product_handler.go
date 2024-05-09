package handlers

import (
	"DB_Project/internal/business/domains/productdomain"
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"encoding/json"
	"log"
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

// GetProduct gets a single product based on its id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := dependencies.Dependencies.ProductDeps.ProductDomain.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, product)
}

// PostProduct posts a product to the database
func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product productdomain.Product

	// Log the start of the function
	log.Println("Starting PostProduct")

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Error decoding product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the product received
	log.Printf("Decoded product: %+v", product)

	// Call the domain function to post the product
	id, err := dependencies.Dependencies.ProductDeps.ProductDomain.PostProduct(&product)
	if err != nil {
		log.Printf("Error posting product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the ID of the product posted
	log.Printf("Product posted with ID: %v", id)

	// Set the product ID
	product.ID = id

	// Log the final product being returned
	log.Printf("Final product to return: %+v", product)

	// Send the response
	utils.JSON(w, http.StatusCreated, product)
}
