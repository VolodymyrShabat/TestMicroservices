package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/pkg/proto" // gRPC generated code
	"net/http"
	"time"
)

type AuthHandlers struct {
	AuthClient authpb.AuthServiceClient
}

// LoginRequest is what we expect from the client
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles /login requests by forwarding them to the Auth microservice.
func (h *AuthHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("trying to login")
	resp, err := h.AuthClient.Login(ctx, &authpb.LoginRequest{
		Username: reqData.Username,
		Password: reqData.Password,
	})
	fmt.Println("trying to login2", err)
	if err != nil {
		// You can map different error types to different status codes if you like
		http.Error(w, fmt.Sprintf("Auth service error: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": resp.GetToken(),
	})
}

// CheckToken handles /sign_in requests by forwarding them to the Auth microservice.
func (h *AuthHandlers) CheckToken(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the JSON request body into a local struct
	var reqData struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Call the Auth microservice via gRPC
	resp, err := h.AuthClient.ValidateToken(ctx, &authpb.ValidateTokenRequest{
		Token: reqData.Token,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Auth service error: %v", err), http.StatusUnauthorized)
		return
	}

	// 4. Write a JSON response containing the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{
		"IsValid": resp.GetIsValid(),
	})
}
