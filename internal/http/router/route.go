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

	// Swagger endpoint (here a client can test the different endpoints and see expected results)
	mux.Handle("GET /electromart/v1/swagger/*", swagger.Handler(swagger.URL("/electromart/v1/swagger/doc.json")))

	// Customers endpoints
	mux.HandleFunc("GET /electromart/v1/customers", middlewares.AdminMiddleware(handlers.AllCustomers))
	mux.HandleFunc("GET /electromart/v1/customers/", middlewares.AdminMiddleware(handlers.AllCustomers))

	mux.HandleFunc("POST /electromart/v1/customers/signup", handlers.Signup)
	mux.HandleFunc("POST /electromart/v1/customers/signup/", handlers.Signup)
	mux.HandleFunc("POST /electromart/v1/customers/login", handlers.Login)
	mux.HandleFunc("POST /electromart/v1/customers/login/", handlers.Login)
	mux.HandleFunc("POST /electromart/v1/customers/logout", handlers.Logout)
	mux.HandleFunc("POST /electromart/v1/customers/logout/", handlers.Logout)

	mux.HandleFunc("GET /electromart/v1/customers/me", handlers.MyProfile)
	mux.HandleFunc("GET /electromart/v1/customers/me/", handlers.MyProfile)
	mux.HandleFunc("PATCH /electromart/v1/customers/me", handlers.UpdateMyProfile)
	mux.HandleFunc("PATCH /electromart/v1/customers/me/", handlers.UpdateMyProfile)
	mux.HandleFunc("DELETE /electromart/v1/customers/me", handlers.DeleteMyProfile)
	mux.HandleFunc("DELETE /electromart/v1/customers/me/", handlers.DeleteMyProfile)

	mux.HandleFunc("GET /electromart/v1/customers/top/{limit}", middlewares.AdminMiddleware(handlers.TopCustomers))

	// Products endpoint
	mux.HandleFunc("GET /electromart/v1/products", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("html") == "true" {
			http.ServeFile(w, r, "public/html/products.html")
		} else {
			handlers.GetAllProducts(w, r)
		}
	})
	mux.HandleFunc("GET /electromart/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("html") == "true" {
			http.ServeFile(w, r, "public/html/products.html")
		} else {
			handlers.GetAllProducts(w, r)
		}
	})

	mux.HandleFunc("GET /electromart/v1/products/{id}", handlers.GetProduct)
	mux.HandleFunc("GET /electromart/v1/products/{id}/", handlers.GetProduct)
	mux.HandleFunc("POST /electromart/v1/products", middlewares.AdminMiddleware(handlers.PostProduct))
	mux.HandleFunc("POST /electromart/v1/products/", middlewares.AdminMiddleware(handlers.PostProduct))
	mux.HandleFunc("PATCH /electromart/v1/products/{id}", middlewares.AdminMiddleware(handlers.PatchProduct))
	mux.HandleFunc("PATCH /electromart/v1/products/{id}/", middlewares.AdminMiddleware(handlers.PatchProduct))
	mux.HandleFunc("DELETE /electromart/v1/products/{id}", middlewares.AdminMiddleware(handlers.DeleteProduct))
	mux.HandleFunc("DELETE /electromart/v1/products/{id}/", middlewares.AdminMiddleware(handlers.DeleteProduct))

	mux.HandleFunc("GET /electromart/v1/products/sales", handlers.TotalSalesPerProduct)
	mux.HandleFunc("GET /electromart/v1/products/sales/", handlers.TotalSalesPerProduct)
	mux.HandleFunc("GET /electromart/v1/products/discounts", handlers.CurrentDiscountedProducts)
	mux.HandleFunc("GET /electromart/v1/products/discounts/", handlers.CurrentDiscountedProducts)
	mux.HandleFunc("GET /electromart/v1/products/full-text-search/{search}", handlers.GetFullTextSearchProduct)
	mux.HandleFunc("GET /electromart/v1/products/full-text-search/{search}/", handlers.GetFullTextSearchProduct)

	// Cart endpoint
	mux.HandleFunc("GET /electromart/v1/cart", handlers.GetCart)
	mux.HandleFunc("GET /electromart/v1/cart/", handlers.GetCart)
	mux.HandleFunc("POST /electromart/v1/cart/{productID}", handlers.AddToCart)
	mux.HandleFunc("POST /electromart/v1/cart/{productID}/", handlers.AddToCart)

	// Checkout endpoint
	mux.HandleFunc("POST /electromart/v1/checkout", handlers.CreateOrder)
	mux.HandleFunc("POST /electromart/v1/checkout/", handlers.CreateOrder)

	// Orders endpoint

	mux.HandleFunc("GET /electromart/v1/orders/{orderId}/details", middlewares.AdminMiddleware(handlers.OrderWithDetails))
	mux.HandleFunc("GET /electromart/v1/orders/{orderId}/details/", middlewares.AdminMiddleware(handlers.OrderWithDetails))

	return mux
}
