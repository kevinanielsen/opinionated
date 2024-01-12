package models

import "gorm.io/gorm"

type Choice struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Question struct {
	gorm.Model
	Title   string   `json:"title"`
	Choices []Choice `json:"choices"`
	Poll    Poll     `json:"poll"`
}
