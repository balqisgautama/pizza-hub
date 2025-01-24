package handlers

import (
	"encoding/json"
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"
)

type MenuHandler struct {
	MenuService *services.MenuService
}

func NewMenuHandler(menuService *services.MenuService) *MenuHandler {
	return &MenuHandler{
		MenuService: menuService,
	}
}

func (h MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMenu := h.MenuService.AddMenu(menu.Name, menu.Duration)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMenu)
}

func (h MenuHandler) GetMenus(w http.ResponseWriter, r *http.Request) {
	menus := h.MenuService.GetMenus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}
