package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter initializes all routes for the API Gateway.
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/login", LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/validate", ValidateTokenHandler).Methods(http.MethodPost)

	// Resource routes
	r.HandleFunc("/resources", ListResourcesHandler).Methods(http.MethodGet)
	r.HandleFunc("/resources/{id}", GetResourceHandler).Methods(http.MethodGet)
	r.HandleFunc("/resources", CreateResourceHandler).Methods(http.MethodPost)

	r.Use(LoggingMiddleware)
	// Additional routes, e.g. for user profiles, payments, etc.

	return r
}
