package handlers

import (
	"DB_Project/internal/business/domains/productdomain"
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

// GetFullTextSearchProduct gets a slice of products based on a string, compared to product description.
func GetFullTextSearchProduct(w http.ResponseWriter, r *http.Request) {
	encodedSearch := r.PathValue("search")

	// decode the search string
	search, err := url.QueryUnescape(encodedSearch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := dependencies.Dependencies.ProductDeps.ProductDomain.SearchProductFullText(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, product)
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

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Error decoding product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the domain function to post the product
	id, err := dependencies.Dependencies.ProductDeps.ProductDomain.PostProduct(&product)
	if err != nil {
		log.Printf("Error posting product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the product ID
	product.ID = id

	// Send the response
	utils.JSON(w, http.StatusCreated, product)
}

// PatchProduct updates a product in the database
func PatchProduct(w http.ResponseWriter, r *http.Request) {
	var product productdomain.Product

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Error decoding product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the product ID from the path
	id := r.PathValue("id")

	// Call the domain function to patch the product
	newProduct, err := dependencies.Dependencies.ProductDeps.ProductDomain.PatchProduct(id, &product)
	if err != nil {
		log.Printf("Error patching product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	utils.JSON(w, http.StatusOK, newProduct)
}

// DeleteProduct deletes a product from the database
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the path
	id := r.PathValue("id")

	// Call the domain function to delete the product
	err := dependencies.Dependencies.ProductDeps.ProductDomain.DeleteProduct(id)
	if err != nil {
		log.Printf("Error deleting product: %v", err) // Log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
