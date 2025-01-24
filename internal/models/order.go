package models

type Order struct {
	ID     int    `json:"id"`
	MenuID int    `json:"menu_id"`
	ChefID int    `json:"chef_id"`
	Status string `json:"status"`
}
