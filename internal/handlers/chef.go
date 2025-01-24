package handlers

import (
	"encoding/json"
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"
)

type ChefHandler struct {
	ChefService *services.ChefService
}

func NewChefHandler(chefService *services.ChefService) *ChefHandler {
	return &ChefHandler{
		ChefService: chefService,
	}
}

func (h ChefHandler) CreateChef(w http.ResponseWriter, r *http.Request) {
	var chef models.Chef
	if err := json.NewDecoder(r.Body).Decode(&chef); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newChef := h.ChefService.AddChef(chef.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newChef)
}

func (h ChefHandler) GetChefs(w http.ResponseWriter, r *http.Request) {
	chefs := h.ChefService.GetChefs()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chefs)
}
