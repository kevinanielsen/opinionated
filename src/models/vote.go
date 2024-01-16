package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID     uint `json:"userId"`
	QuestionID uint `json:"questionId"`
}
