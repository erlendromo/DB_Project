package middlewares

import (
	"DB_Project/internal/constants"
	"fmt"
	"net/http"
	"time"
)

type Logger struct {
	Handler http.Handler
}

func NewLogger(h http.Handler) http.Handler {
	return &Logger{
		Handler: h,
	}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.Handler.ServeHTTP(w, r)

	executionTime := fmt.Sprintf("%dμs", time.Since(start).Microseconds())
	timeUTC := fmt.Sprintf("%s UTC", time.Now().UTC().Format(constants.TIME_FORMAT))

	fmt.Printf("| %s | %s | %s | %s |\n", timeUTC, r.Method, r.URL.Path, executionTime)
}
