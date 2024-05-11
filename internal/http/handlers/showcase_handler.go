package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"net/http"
)

// TopCustomers handles the request to identify top customers based on the limit specified
func TopCustomers(w http.ResponseWriter, r *http.Request) {
	limit := r.PathValue("limit")

	topCustomers, err := dependencies.Dependencies.ShowcaseDeps.ShowcaseDomain.IdentifyTopCustomers(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, topCustomers)
}

// OrderWithDetails handles the request to fetch orders along with their detailed information
func OrderWithDetails(w http.ResponseWriter, r *http.Request) {

	orderId := r.PathValue("orderId")

	orderDetails, err := dependencies.Dependencies.ShowcaseDeps.ShowcaseDomain.FetchOrderWithDetails(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, orderDetails)
}

// CurrentDiscountedProducts handles the request to list currently discounted products
func CurrentDiscountedProducts(w http.ResponseWriter, r *http.Request) {
	discountedProducts, err := dependencies.Dependencies.ShowcaseDeps.ShowcaseDomain.ListCurrentDiscountedProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, discountedProducts)
}

// TotalSalesPerProduct handles the request for calculating total sales per product
func TotalSalesPerProduct(w http.ResponseWriter, r *http.Request) {
	sales, err := dependencies.Dependencies.ShowcaseDeps.ShowcaseDomain.CalculateTotalSalesPerProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, sales)
}
