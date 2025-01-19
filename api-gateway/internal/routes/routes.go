package routes

import (
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/handlers"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter initializes all routes for the API Gateway.
func SetupRouter(ah handlers.AuthHandlers, rh handlers.ResourceHandlers) *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/users/sign_in", ah.LoginHandler).Methods(http.MethodPost)

	// Resource routes
	r.HandleFunc("/users", rh.GetUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/books", rh.GetBooksHandler).Methods(http.MethodGet)

	r.Use(middlewares.LoggingMiddleware)
	// Additional routes, e.g. for user profiles, payments, etc.

	return r
}
