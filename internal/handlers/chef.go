package handlers

import (
	"net/http"

	"pizza-hub/internal/models"
	"pizza-hub/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type ChefHandler struct {
	Validate    *validator.Validate
	ChefService *services.ChefService
}

func NewChefHandler(validate *validator.Validate, chefService *services.ChefService) *ChefHandler {
	return &ChefHandler{
		Validate:    validate,
		ChefService: chefService,
	}
}

func (h ChefHandler) CreateChef(c echo.Context) error {
	var chef models.Chef
	if err := c.Bind(&chef); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.Validate.Struct(chef); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newChef := h.ChefService.AddChef(chef.Name)

	return c.JSON(http.StatusCreated, newChef)
}

func (h ChefHandler) GetChefs(c echo.Context) error {
	chefs := h.ChefService.GetChefs()
	return c.JSON(http.StatusOK, chefs)
}
