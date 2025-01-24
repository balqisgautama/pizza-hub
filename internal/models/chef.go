package models

type Chef struct {
	ID     int    `json:"id"`
	Name   string `json:"name" validate:"required,min=3,max=100"`
	Status string `json:"status"`
}
