package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/services"
	"net/http"
)

type ResourcesHandlers struct {
	ResourcesService *services.ResourcesService
}

func NewResourcesHandlers(resourcesService *services.ResourcesService) *ResourcesHandlers {
	return &ResourcesHandlers{
		ResourcesService: resourcesService,
	}
}

func (h *ResourcesHandlers) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.ResourcesService.GetUsers()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ResourcesHandlers) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.ResourcesService.GetBooks()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
