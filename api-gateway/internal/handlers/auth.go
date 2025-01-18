package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/proto/auth" // gRPC generated code
	"net/http"
	"time"
)

// LoginRequest is what we expect from the client
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles /login requests by forwarding them to the Auth microservice.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Parse request
	var reqData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Call Auth microservice via gRPC
	// Assume you have a global or injected gRPC client, e.g., `authClient authpb.AuthServiceClient`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := authClient.Login(ctx, &authpb.AuthRequest{
		Username: reqData.Username,
		Password: reqData.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Auth service error: %v", err), http.StatusUnauthorized)
		return
	}

	// 3. Respond to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": resp.GetToken(),
	})
}
