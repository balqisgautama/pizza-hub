package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pizza-hub/internal/chef"
	"pizza-hub/internal/models"
	"pizza-hub/menu"
	"sync"
)

type PizzaHubAPI struct {
	Chefs  []*chef.Chef
	Orders []*models.Order
	Menu   *menu.Menu
	mu     sync.Mutex
}

func NewPizzaHubAPI() *PizzaHubAPI {
	return &PizzaHubAPI{
		Chefs:  []*chef.Chef{},
		Orders: []*models.Order{},
		Menu:   menu.NewMenu(),
	}
}

func (ph *PizzaHubAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		switch r.URL.Path {
		case "/chefs":
			ph.AddChef()
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Chef added\n")
		case "/menus":
			var menuItem menu.MenuItem
			if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ph.Menu.AddMenuItem(menuItem)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Menu item added\n")
		case "/orders":
			var order models.Order
			if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := ph.Menu.GetMenu()[order.Pizza]; !ok {
				http.Error(w, "Invalid pizza type", http.StatusBadRequest)
				return
			}
			if ph.AssignOrder(&order) {
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, "Order %d assigned to chef\n", order.ID)
			} else {
				http.Error(w, "No available chef", http.StatusServiceUnavailable)
			}
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	case http.MethodGet:
		switch r.URL.Path {
		case "/chefs":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ph.Chefs)
		case "/menus":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ph.Menu.GetMenu())
		case "/orders":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ph.Orders)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *PizzaHubAPI) AddChef() *chef.Chef {
	ph.mu.Lock()
	defer ph.mu.Unlock()
	chef := chef.NewChef(len(ph.Chefs) + 1)
	ph.Chefs = append(ph.Chefs, chef)
	return chef
}

func (ph *PizzaHubAPI) AssignOrder(order *models.Order) bool {
	ph.mu.Lock()
	defer ph.mu.Unlock()
	if len(ph.Chefs) == 0 {
		return false
	}
	for _, chef := range ph.Chefs {
		if chef.IsAvailable() {
			chef.AssignOrder(order)
			order.ID = len(ph.Orders) + 1
			ph.Orders = append(ph.Orders, order)
			return true
		}
	}
	return false
}
