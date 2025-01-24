package menu

import (
	"testing"
)

func TestAddMenuItem(t *testing.T) {
	menu := NewMenu()
	menu.AddMenuItem(MenuItem{Name: "Pepperoni", Duration: 4})

	expected := map[string]int{"Cheese": 3, "BBQ": 5, "Pepperoni": 4}
	result := menu.GetMenu()

	if len(result) != len(expected) {
		t.Errorf("Expected menu length %d, got %d", len(expected), len(result))
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected menu item %s with duration %d, got %d", k, v, result[k])
		}
	}
}

func TestGetMenu(t *testing.T) {
	menu := NewMenu()
	expected := map[string]int{"Cheese": 3, "BBQ": 5}
	result := menu.GetMenu()

	if len(result) != len(expected) {
		t.Errorf("Expected menu length %d, got %d", len(expected), len(result))
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected menu item %s with duration %d, got %d", k, v, result[k])
		}
	}
}
