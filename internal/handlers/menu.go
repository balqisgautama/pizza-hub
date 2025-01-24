package handlers

import (
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type MenuHandler struct {
	Validate    *validator.Validate
	MenuService *services.MenuService
}

func NewMenuHandler(validate *validator.Validate, menuService *services.MenuService) *MenuHandler {
	return &MenuHandler{
		Validate:    validate,
		MenuService: menuService,
	}
}

func (h MenuHandler) CreateMenu(c echo.Context) error {
	var menu models.Menu
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.Validate.Struct(menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newMenu := h.MenuService.AddMenu(menu.Name, menu.Duration)
	return c.JSON(http.StatusCreated, newMenu)
}

func (h MenuHandler) GetMenus(c echo.Context) error {
	menus := h.MenuService.GetMenus()
	return c.JSON(http.StatusOK, menus)
}
