package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Name      string     `json:"name"`
	UserID    uint       `json:"userId"`
	Questions []Question `json:"questions"`
}
