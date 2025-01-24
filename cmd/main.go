package main

import (
	"pizza-hub/internal/handlers"
	"pizza-hub/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	validate := validator.New()

	chefService := services.NewChefService()
	menuService := services.NewMenuService()
	orderService := services.NewOrderService(chefService, menuService)

	// Chef endpoints
	chefHandler := handlers.NewChefHandler(validate, chefService)
	e.POST("/chefs", chefHandler.CreateChef)
	e.GET("/chefs", chefHandler.GetChefs)

	// Menu endpoints
	menuHandler := handlers.NewMenuHandler(validate, menuService)
	e.POST("/menus", menuHandler.CreateMenu)
	e.GET("/menus", menuHandler.GetMenus)

	// Order endpoints
	orderHandler := handlers.NewOrderHandler(validate, chefService, menuService, orderService)
	e.POST("/orders", orderHandler.CreateOrder)
	e.GET("/orders", orderHandler.GetOrders)

	e.Logger.Fatal(e.Start(":8080"))
}
