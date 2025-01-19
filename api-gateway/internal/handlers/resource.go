package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	resourcepb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
	"net/http"
	"time"
)

type ResourceHandlers struct {
	ResourceClient resourcepb.ResourceServiceClient
}

// LoginHandler handles /login requests by forwarding them to the Auth microservice.
func (h *ResourceHandlers) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.ResourceClient.GetUsers(ctx, &resourcepb.EmptyRequest{})
	if err != nil {
		// You can map different error types to different status codes if you like
		http.Error(w, fmt.Sprintf("Resource service error: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": resp.GetUsers(),
	})
}

// LoginHandler handles /login requests by forwarding them to the Auth microservice.
func (h *ResourceHandlers) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.ResourceClient.GetBooks(ctx, &resourcepb.EmptyRequest{})
	if err != nil {
		// You can map different error types to different status codes if you like
		http.Error(w, fmt.Sprintf("Resource service error: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"books": resp.GetBooks(),
	})
}
