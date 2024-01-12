package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	User     User     `json:"user"`
	Question Question `json:"question"`
}
