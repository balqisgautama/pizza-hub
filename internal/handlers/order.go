package handlers

import (
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type OrderHandler struct {
	Validate     *validator.Validate
	ChefService  *services.ChefService
	MenuService  *services.MenuService
	OrderService *services.OrderService
}

func NewOrderHandler(
	validate *validator.Validate,
	chefService *services.ChefService,
	menuService *services.MenuService,
	orderService *services.OrderService,
) *OrderHandler {
	return &OrderHandler{
		Validate:     validate,
		ChefService:  chefService,
		MenuService:  menuService,
		OrderService: orderService,
	}
}

func (h OrderHandler) CreateOrder(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.Validate.Struct(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newOrder, err := h.OrderService.AddOrder(order.MenuID, order.ChefID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newOrder)
}

func (h OrderHandler) GetOrders(c echo.Context) error {
	orders := h.OrderService.GetOrders()
	return c.JSON(http.StatusOK, orders)
}
