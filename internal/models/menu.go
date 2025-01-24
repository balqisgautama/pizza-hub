package models

type Menu struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Duration int    `json:"duration" validate:"required,min=1"`
}
