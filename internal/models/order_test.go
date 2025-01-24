package models

import (
	"testing"
)

func TestOrderStruct(t *testing.T) {
	order := Order{ID: 1, Pizza: "Cheese", ChefID: 1}
	if order.ID != 1 {
		t.Errorf("Expected order ID 1, got %d", order.ID)
	}
	if order.Pizza != "Cheese" {
		t.Errorf("Expected order Pizza 'Cheese', got %s", order.Pizza)
	}
	if order.ChefID != 1 {
		t.Errorf("Expected order ChefID 1, got %d", order.ChefID)
	}
}
