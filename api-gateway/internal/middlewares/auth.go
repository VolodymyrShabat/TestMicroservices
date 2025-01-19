package middlewares

import (
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

func NewAuthMiddleware(AuthService *services.AuthService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			req := &models.IsTokenValidRequest{
				Token: token,
			}

			res, err := AuthService.ValidateToken(req.Token)
			if err != nil {
				http.Error(w, "Non valid auth token", http.StatusInternalServerError)
				return
			}
			if res.IsTokenValid {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		})
	}
}
