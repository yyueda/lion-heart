package router

import (
	"net/http"

	"github.com/yyueda/lion-heart/backend/user-service/internal/handler"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()

	// router.HandleFunc("/health", handler.HealthCheck)
	router.HandleFunc("/users", handler.GetUsers)

	return loggingMiddleware(router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("[LOG] Request (user-service):", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
