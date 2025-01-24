package menu

import (
	"sync"
)

type MenuItem struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type Menu struct {
	mu    sync.Mutex
	menus map[string]int
}

func NewMenu() *Menu {
	return &Menu{
		menus: map[string]int{
			"Cheese": 3, // Duration in seconds
			"BBQ":    5,
		},
	}
}

func (m *Menu) AddMenuItem(item MenuItem) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.menus[item.Name] = item.Duration
}

func (m *Menu) GetMenu() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.menus
}
