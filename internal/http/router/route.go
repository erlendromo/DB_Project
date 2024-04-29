package router

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// TODO setup paths for endpoints

	return mux
}
