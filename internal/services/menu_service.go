package services

import (
	"errors"
	"sync"

	"pizza-hub/internal/models"
)

type MenuService struct {
	menus  []*models.Menu
	nextID int
	mu     sync.Mutex
}

func NewMenuService() *MenuService {
	return &MenuService{
		menus:  []*models.Menu{},
		nextID: 1,
		mu:     sync.Mutex{},
	}
}

func (s *MenuService) AddMenu(name string, duration int) *models.Menu {
	s.mu.Lock()
	defer s.mu.Unlock()

	menu := &models.Menu{
		ID:       s.nextID,
		Name:     name,
		Duration: duration,
	}
	s.menus = append(s.menus, menu)
	s.nextID++
	return menu
}

func (s *MenuService) GetMenus() []*models.Menu {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.menus
}

func (s *MenuService) FindMenuByID(id int) (*models.Menu, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, menu := range s.menus {
		if menu.ID == id {
			return menu, nil
		}
	}

	return nil, errors.New("menu not found")
}
