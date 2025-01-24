package chef

import (
	"log"
	"pizza-hub/internal/models"
	"pizza-hub/menu"
	"sync"
	"time"
)

type Chef struct {
	ID     int
	IsBusy bool
	mu     sync.Mutex
	orders chan *models.Order
}

func NewChef(id int) *Chef {
	return &Chef{
		ID:     id,
		IsBusy: false,
		orders: make(chan *models.Order, 1),
	}
}

func (c *Chef) Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	menus := menu.NewMenu()
	for order := range c.orders {
		c.mu.Lock()
		c.IsBusy = true
		c.mu.Unlock()

		log.Printf("Chef %d is processing order %d for %s pizza\n", c.ID, order.ID, order.Pizza)
		time.Sleep(time.Duration(menus.GetMenu()[order.Pizza]) * time.Second)
		log.Printf("Chef %d finished processing order %d for %s pizza\n", c.ID, order.ID, order.Pizza)

		c.mu.Lock()
		c.IsBusy = false
		c.mu.Unlock()
	}
}

func (c *Chef) AssignOrder(order *models.Order) {
	c.orders <- order
}

func (c *Chef) IsAvailable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return !c.IsBusy
}
