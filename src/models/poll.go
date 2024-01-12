package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Name      string     `json:"name"`
	User      User       `json:"user"`
	Questions []Question `json:"questions"`
}
