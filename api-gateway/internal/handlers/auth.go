package handlers

import (
	"encoding/json"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/services"
	"net/http"
)

type AuthHandlers struct {
	AuthService *services.AuthService
}

func NewAuthHandlers(authService *services.AuthService) *AuthHandlers {
	return &AuthHandlers{
		AuthService: authService,
	}
}

func (h *AuthHandlers) SignInHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.SignInRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.SignIn(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Error during call", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, "Error during sending response", http.StatusInternalServerError)
		return
	}
}
