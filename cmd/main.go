package main

import (
	"log"
	"net/http"
	"pizza-hub/internal/api"
	"sync"
)

func main() {
	ph := api.NewPizzaHubAPI()
	var wg sync.WaitGroup

	http.Handle("/", ph)

	// Start serving chefs
	for _, chef := range ph.Chefs {
		wg.Add(1)
		go chef.Serve(&wg)
	}

	log.Println("PizzaHub is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	wg.Wait()
}
