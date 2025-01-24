package models

type Order struct {
	ID     int    `json:"id"`
	Pizza  string `json:"pizza"`
	ChefID int    `json:"chef_id"`
}
