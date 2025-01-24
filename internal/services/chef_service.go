package services

import (
	"errors"
	"sync"

	"pizza-hub/internal/models"
)

type ChefService struct {
	chefs  []*models.Chef
	nextID int
	mu     sync.Mutex
}

func NewChefService() *ChefService {
	return &ChefService{
		chefs:  []*models.Chef{},
		nextID: 1,
		mu:     sync.Mutex{},
	}
}

func (s *ChefService) AddChef(name string) *models.Chef {
	s.mu.Lock()
	defer s.mu.Unlock()

	chef := &models.Chef{
		ID:     s.nextID,
		Name:   name,
		Status: "available",
	}
	s.chefs = append(s.chefs, chef)
	s.nextID++
	return chef
}

func (s *ChefService) GetChefs() []*models.Chef {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.chefs
}

func (s *ChefService) FindChefByID(id int) (*models.Chef, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, chef := range s.chefs {
		if chef.ID == id {
			return chef, nil
		}
	}

	return nil, errors.New("chef not found")
}

// New function to find chefs by status
func (s *ChefService) FindChefsByStatus(status string) ([]*models.Chef, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var filteredChefs []*models.Chef
	for _, chef := range s.chefs {
		if chef.Status == status {
			filteredChefs = append(filteredChefs, chef)
		}
	}

	if len(filteredChefs) == 0 {
		return s.chefs, errors.New("no chefs available")
	}

	return filteredChefs, nil
}
