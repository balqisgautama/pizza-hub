package services

import (
	"errors"
	"sync"
	"time"

	"pizza-hub/internal/models"
)

type OrderService struct {
	orders []*models.Order
	nextID int
	chefs  *ChefService
	menus  *MenuService
	mu     sync.Mutex
}

func NewOrderService(chefs *ChefService, menus *MenuService) *OrderService {
	return &OrderService{
		orders: []*models.Order{},
		nextID: 1,
		chefs:  chefs,
		menus:  menus,
		mu:     sync.Mutex{},
	}
}

func (s *OrderService) AddOrder(menuID, chefID int) (*models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.chefs.GetChefs()) == 0 || len(s.menus.GetMenus()) == 0 {
		return nil, errors.New("no chefs or menus available")
	}

	menu, err := s.menus.FindMenuByID(menuID)
	if err != nil {
		return nil, err
	}

	chef, err := s.chefs.FindChefByID(chefID)
	if err != nil {
		chefs, err := s.chefs.FindChefsByStatus("available")
		if err != nil {
			return nil, err
		}
		chef = chefs[0]
	}

	if condition := chef.Status == "busy"; condition {
		return nil, errors.New("chef is busy")
	}

	order := &models.Order{
		ID:     s.nextID,
		MenuID: menuID,
		ChefID: chef.ID,
		Status: "processing",
	}
	s.orders = append(s.orders, order)
	s.nextID++

	go s.processOrder(order, menu, chef)
	return order, nil
}

func (s *OrderService) GetOrders() []*models.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.orders
}

func (s *OrderService) processOrder(order *models.Order, menu *models.Menu, chef *models.Chef) {
	chef.Status = "busy"
	time.Sleep(time.Duration(menu.Duration) * time.Second)

	s.mu.Lock()
	defer s.mu.Unlock()

	order.Status = "completed"
	chef.Status = "available"
}
