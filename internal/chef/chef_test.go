package chef

import (
	"pizza-hub/internal/models"
	"sync"
	"testing"
	"time"
)

func TestChefServe(t *testing.T) {
	chef := NewChef(1)
	var wg sync.WaitGroup
	wg.Add(1)
	go chef.Serve(&wg)

	order := &models.Order{ID: 1, Pizza: "Cheese"}
	chef.AssignOrder(order)

	time.Sleep(4 * time.Second) // Wait for the order to be processed

	if chef.IsAvailable() != true {
		t.Errorf("Expected chef to be available, got busy")
	}
}

func TestChefAssignOrder(t *testing.T) {
	chef := NewChef(1)
	order := &models.Order{ID: 1, Pizza: "Cheese"}
	chef.AssignOrder(order)

	if chef.IsAvailable() != false {
		t.Errorf("Expected chef to be busy, got available")
	}
}

func TestChefIsAvailable(t *testing.T) {
	chef := NewChef(1)
	if chef.IsAvailable() != true {
		t.Errorf("Expected chef to be available, got busy")
	}

	order := &models.Order{ID: 1, Pizza: "Cheese"}
	chef.AssignOrder(order)

	if chef.IsAvailable() != false {
		t.Errorf("Expected chef to be busy, got available")
	}
}
