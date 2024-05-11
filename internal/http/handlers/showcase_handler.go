package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"net/http"
)

// TopCustomers handles the request to identify top customers based on the limit specified
//
//	@title			TopCustomers
//	@summary		Identify top customers
//	@description	Identify top customers (requires admin login)
//	@tags			Showcase
//	@security		AdminAuth
//	@produce		json
//	@param			limit	path		string	true	"Limit"
//	@success		200		{array}		showcasedomain.TopCustomer
//	@failure		401		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/top/{limit} [get]
func TopCustomers(w http.ResponseWriter, r *http.Request) {
	limit := r.PathValue("limit")

	topCustomers, err := dependencies.Dependencies.ShowcaseDeps.PSQLShowcase.IdentifyTopCustomers(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, topCustomers)
}

// OrderWithDetails handles the request to fetch orders along with their detailed information
//
//	@title			OrderWithDetails
//	@summary		Fetch order with details
//	@description	Fetch order with details (requires admin login)
//	@tags			Showcase
//	@security		AdminAuth
//	@produce		json
//	@param			orderID	path		string	true	"Order ID"
//	@success		200		{object}	showcasedomain.OrderDetail
//	@failure		401		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/orders/{orderID}/details [get]
func OrderWithDetails(w http.ResponseWriter, r *http.Request) {

	orderId := r.PathValue("orderID")

	orderDetails, err := dependencies.Dependencies.ShowcaseDeps.PSQLShowcase.FetchOrderWithDetails(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, orderDetails)
}

// CurrentDiscountedProducts handles the request to list currently discounted products
//
//	@title			CurrentDiscountedProducts
//	@summary		List current discounted products
//	@description	List current discounted products
//	@tags			Showcase
//	@produce		json
//	@success		200	{array}		showcasedomain.DiscountedProduct
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/discounts [get]
func CurrentDiscountedProducts(w http.ResponseWriter, r *http.Request) {
	discountedProducts, err := dependencies.Dependencies.ShowcaseDeps.PSQLShowcase.ListCurrentDiscountedProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, discountedProducts)
}

// TotalSalesPerProduct handles the request for calculating total sales per product
//
//	@title			TotalSalesPerProduct
//	@summary		Calculate total sales per product
//	@description	Calculate total sales per product
//	@tags			Showcase
//	@produce		json
//	@success		200	{array}		showcasedomain.ProductSales
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/products/sales-per-product [get]
func TotalSalesPerProduct(w http.ResponseWriter, r *http.Request) {
	sales, err := dependencies.Dependencies.ShowcaseDeps.PSQLShowcase.CalculateTotalSalesPerProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, sales)
}
