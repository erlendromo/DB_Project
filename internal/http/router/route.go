package router

import (
	"net/http"

	_ "DB_Project/docs"
	"DB_Project/internal/http/handlers"
	"DB_Project/internal/http/middlewares"

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

	// mux.HandleFunc("GET /electromart/v1/myprofile/", handlers.GetUserByID)
	// mux.HandleFunc("GET /electromart/v1/myprofile/", handlers.GetUserByID)

	return mux
}
