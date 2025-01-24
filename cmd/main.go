package main

import (
	"log"
	"net/http"

	"pizza-hub/internal/handlers"
	"pizza-hub/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	chefService := services.NewChefService()
	menuService := services.NewMenuService()
	orderService := services.NewOrderService(chefService, menuService)

	// Chef endpoints
	chefHandler := handlers.NewChefHandler(chefService)
	r.HandleFunc("/chefs", chefHandler.CreateChef).Methods("POST")
	r.HandleFunc("/chefs", chefHandler.GetChefs).Methods("GET")

	// Menu endpoints
	menuHandler := handlers.NewMenuHandler(menuService)
	r.HandleFunc("/menus", menuHandler.CreateMenu).Methods("POST")
	r.HandleFunc("/menus", menuHandler.GetMenus).Methods("GET")

	// Order endpoints
	orderHandler := handlers.NewOrderHandler(chefService, menuService, orderService)
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")

	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
