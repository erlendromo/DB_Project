package router

import (
	_ "DB_Project/docs"
	"DB_Project/internal/http/handlers"
	"DB_Project/internal/http/middlewares"
	"net/http"

	swagger "github.com/swaggo/http-swagger/v2"
)

// @title			Electromart API
// @version		1.0
// @description	This is a json RESTful API for the newly established e-commerce Electromart
// @termsOfService	http://swagger.io/terms/
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// TODO setup paths for endpoints
	mux.Handle("GET /electromart/v1/swagger/*", swagger.Handler(swagger.URL("/electromart/v1/swagger/doc.json")))

	mux.HandleFunc("POST /electromart/v1/signup", handlers.Signup)
	mux.HandleFunc("POST /electromart/v1/signup/", handlers.Signup)
	mux.HandleFunc("POST /electromart/v1/login", handlers.Login)
	mux.HandleFunc("POST /electromart/v1/login/", handlers.Login)
	mux.HandleFunc("POST /electromart/v1/logout", handlers.Logout)
	mux.HandleFunc("POST /electromart/v1/logout/", handlers.Logout)

	mux.HandleFunc("GET /electromart/v1/myprofile", handlers.MyProfile)
	mux.HandleFunc("GET /electromart/v1/myprofile/", handlers.MyProfile)

	mux.HandleFunc("GET /electromart/v1/customers", middlewares.AdminMiddleware(handlers.AllCustomers))
	mux.HandleFunc("GET /electromart/v1/customers/", middlewares.AdminMiddleware(handlers.AllCustomers))

	// Products endpoint
	mux.HandleFunc("GET /electromart/v1/products", handlers.GetAllProducts)
	mux.HandleFunc("GET /electromart/v1/products/", handlers.GetAllProducts)

	mux.HandleFunc("GET /electromart/v1/products/{id}", handlers.GetProduct)

	// full text search, searches on product description, it is case-insensitive and needs to full word exact word match with any word in a description
	mux.HandleFunc("GET /electromart/v1/products/full-text-search/{search}", handlers.GetFullTextSearchProduct)

	mux.HandleFunc("POST /electromart/v1/products", middlewares.AdminMiddleware(handlers.PostProduct))
	mux.HandleFunc("POST /electromart/v1/products/", middlewares.AdminMiddleware(handlers.PostProduct))

	mux.HandleFunc("DELETE /electromart/v1/products/{id}", handlers.DeleteProduct)

	mux.HandleFunc("PATCH /electromart/v1/products/{id}", handlers.PatchProduct)

	// UI
	mux.HandleFunc("/html/product", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/html/products.html")
	})

	// showcase handler - the queries we showcase in the report

	mux.HandleFunc("GET /electromart/v1/orders-details/{orderId}", handlers.OrderWithDetails)

	mux.HandleFunc("GET /electromart/v1/discounted-products", handlers.CurrentDiscountedProducts)

	mux.HandleFunc("GET /electromart/v1/sales-per-product", handlers.TotalSalesPerProduct)

	mux.HandleFunc("GET /electromart/v1/top-customers/{limit}", handlers.TopCustomers)

	// mux.HandleFunc("GET /electromart/v1/myprofile/", handlers.GetUserByID)
	// mux.HandleFunc("GET /electromart/v1/myprofile/", handlers.GetUserByID)

	return mux
}
