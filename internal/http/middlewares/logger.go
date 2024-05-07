package middlewares

import (
	"DB_Project/internal/constants"
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		executionTime := fmt.Sprintf("%dμs", time.Since(start).Microseconds())
		timeUTC := fmt.Sprintf("%s UTC", time.Now().UTC().Format(constants.TIME_FORMAT))

		fmt.Printf("| %s | %s | %s | %s |\n", timeUTC, r.Method, r.URL.Path, executionTime)
	})
}
