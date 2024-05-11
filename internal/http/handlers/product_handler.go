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
//
//	@title			GetAllProducts
//	@summary		Get all products
//	@description	Get all products
//	@tags			Products
//	@produce		json
//	@success		200	{array}		productdomain.Product
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/products [get]
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := dependencies.Dependencies.ProductDeps.PSQLProduct.GetAllProducts()
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, products)
}

// GetFullTextSearchProduct gets a slice of products based on a string, compared to product description.
//
//	@title			GetFullTextSearchProduct
//	@summary		Get products based on full text search
//	@description	Get products based on full text search
//	@tags			Products
//	@produce		json
//	@param			search	path		string	true	"Search string"
//	@success		200		{array}		productdomain.Product
//	@failure		400		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/full-text-search/{search} [get]
func GetFullTextSearchProduct(w http.ResponseWriter, r *http.Request) {
	encodedSearch := r.PathValue("search")

	// decode the search string
	search, err := url.QueryUnescape(encodedSearch)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	products, err := dependencies.Dependencies.ProductDeps.PSQLProduct.SearchProductFullText(search)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, products)
}

// GetProduct gets a single product based on its id
//
//	@title			GetProduct
//	@summary		Get a product
//	@description	Get a product
//	@tags			Products
//	@produce		json
//	@param			id	path		string	true	"Product ID"
//	@success		200	{object}	productdomain.Product
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/{id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := dependencies.Dependencies.ProductDeps.PSQLProduct.GetProduct(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, product)
}

// PostProduct posts a product to the database
//
//	@title			PostProduct
//	@summary		Post a product
//	@description	Post a product (requires admin login)
//	@tags			Products
//	@security		AdminAuth
//	@accept			json
//	@produce		json
//	@param			body	body		productdomain.Product	true	"Product to post"
//	@success		201		{object}	productdomain.Product
//	@failure		400		{object}	utils.ErrorResponse
//	@failure		401		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/products [post]
func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product productdomain.Product

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Error decoding product: %v", err) // Log the error
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Call the domain function to post the product
	id, err := dependencies.Dependencies.ProductDeps.PSQLProduct.PostProduct(&product)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Set the product ID
	product.ID = id

	// Send the response
	utils.JSON(w, http.StatusCreated, product)
}

// PatchProduct updates a product in the database
//
//	@title			PatchProduct
//	@summary		Patch a product
//	@description	Patch a product (requires admin login)
//	@tags			Products
//	@security		AdminAuth
//	@accept			json
//	@produce		json
//	@param			id		path		string							true	"Product ID"
//	@param			body	body		productdomain.PointerProduct	true	"Product to patch"
//	@success		200		{object}	productdomain.Product
//	@failure		400		{object}	utils.ErrorResponse
//	@failure		401		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/{id} [patch]
func PatchProduct(w http.ResponseWriter, r *http.Request) {
	var product productdomain.PointerProduct

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Error decoding product: %v", err) // Log the error
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Get the product ID from the path
	id := r.PathValue("id")

	// Call the domain function to patch the product
	newProduct, err := dependencies.Dependencies.ProductDeps.PSQLProduct.PatchProduct(id, &product)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Send the response
	utils.JSON(w, http.StatusOK, newProduct)
}

// DeleteProduct deletes a product from the database
//
//	@title			DeleteProduct
//	@summary		Delete a product
//	@description	Delete a product (requires admin login)
//	@tags			Products
//	@security		AdminAuth
//	@param			id	path	string	true	"Product ID"
//	@success		204
//	@failure		401	{object}	utils.ErrorResponse
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the path
	id := r.PathValue("id")

	// Call the domain function to delete the product
	err := dependencies.Dependencies.ProductDeps.PSQLProduct.DeleteProduct(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
