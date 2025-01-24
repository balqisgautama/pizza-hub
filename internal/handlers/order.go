package handlers

import (
	"encoding/json"
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"
)

type OrderHandler struct {
	ChefService  *services.ChefService
	MenuService  *services.MenuService
	OrderService *services.OrderService
}

func NewOrderHandler(chefService *services.ChefService, menuService *services.MenuService, orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		ChefService:  chefService,
		MenuService:  menuService,
		OrderService: orderService,
	}
}

func (h OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newOrder, err := h.OrderService.AddOrder(order.MenuID, order.ChefID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrder)
}

func (h OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.OrderService.GetOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
