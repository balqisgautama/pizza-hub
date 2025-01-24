package models

type Order struct {
	ID     int    `json:"id"`
	MenuID int    `json:"menu_id" validate:"required,min=1"`
	ChefID int    `json:"chef_id" validate:"omitempty,min=1"`
	Status string `json:"status"`
}
