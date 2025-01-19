package routes

import (
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/handlers"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter initializes all routes for the API Gateway.
func SetupRouter(ah *handlers.AuthHandlers, rh *handlers.ResourcesHandlers) *mux.Router {
	r := mux.NewRouter()
	nonAuthenicatedApi := r.PathPrefix("/api/v1").Subrouter()
	authenticatedApi := r.PathPrefix("/api/v1").Subrouter()
	// Auth routes
	nonAuthenicatedApi.HandleFunc("/users/sign_in", ah.SignInHandler).Methods(http.MethodPost)
	// Resource routes
	nonAuthenicatedApi.HandleFunc("/books", rh.GetBooksHandler).Methods(http.MethodGet)
	authenticatedApi.HandleFunc("/users", rh.GetUsersHandler).Methods(http.MethodGet)
	authenticatedApi.Use(middlewares.NewAuthMiddleware(ah.AuthService))

	r.Use(middlewares.LoggingMiddleware)

	return r
}
