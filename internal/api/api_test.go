package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pizza-hub/internal/chef"
	"pizza-hub/internal/models"
	"pizza-hub/menu"
	"testing"
)

func TestAddChef(t *testing.T) {
	ph := NewPizzaHubAPI()
	req, _ := http.NewRequest("POST", "/chefs", nil)
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
	if w.Body.String() != "Chef added\n" {
		t.Errorf("Expected body 'Chef added\n', got %s", w.Body.String())
	}
}

func TestGetChefs(t *testing.T) {
	ph := NewPizzaHubAPI()
	ph.AddChef()
	req, _ := http.NewRequest("GET", "/chefs", nil)
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var chefs []*chef.Chef
	json.Unmarshal(w.Body.Bytes(), &chefs)

	if len(chefs) != 1 {
		t.Errorf("Expected 1 chef, got %d", len(chefs))
	}
}

func TestAddMenu(t *testing.T) {
	ph := NewPizzaHubAPI()
	menuItem := menu.MenuItem{Name: "Pepperoni", Duration: 4}
	jsonData, _ := json.Marshal(menuItem)
	req, _ := http.NewRequest("POST", "/menus", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
	if w.Body.String() != "Menu item added\n" {
		t.Errorf("Expected body 'Menu item added\n', got %s", w.Body.String())
	}
}

func TestGetMenus(t *testing.T) {
	ph := NewPizzaHubAPI()
	ph.Menu.AddMenuItem(menu.MenuItem{Name: "Pepperoni", Duration: 4})
	req, _ := http.NewRequest("GET", "/menus", nil)
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expected := map[string]int{"Cheese": 3, "BBQ": 5, "Pepperoni": 4}
	var result map[string]int
	json.Unmarshal(w.Body.Bytes(), &result)

	if len(result) != len(expected) {
		t.Errorf("Expected menu length %d, got %d", len(expected), len(result))
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected menu item %s with duration %d, got %d", k, v, result[k])
		}
	}
}

func TestAddOrder(t *testing.T) {
	ph := NewPizzaHubAPI()
	ph.AddChef()

	order := models.Order{ID: 1, Pizza: "Cheese"}
	jsonData, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
	if w.Body.String() != "Order 1 assigned to chef\n" {
		t.Errorf("Expected body 'Order 1 assigned to chef\n', got %s", w.Body.String())
	}
}

func TestGetOrders(t *testing.T) {
	ph := NewPizzaHubAPI()
	ph.AddChef()

	order := models.Order{ID: 1, Pizza: "Cheese"}
	jsonData, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/orders", nil)
	w = httptest.NewRecorder()
	ph.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var orders []*models.Order
	json.Unmarshal(w.Body.Bytes(), &orders)

	if len(orders) != 1 {
		t.Errorf("Expected 1 order, got %d", len(orders))
	}
}
