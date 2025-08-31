package router

import (
	"net/http"

	"github.com/yyueda/lion-heart/backend/donation-service/internal/handler"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()

	// router.HandleFunc("/health", handler.HealthCheck)
	router.HandleFunc("/donations", handler.GetDonations)

	return loggingMiddleware(router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("[LOG] Request (donation-service):", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
