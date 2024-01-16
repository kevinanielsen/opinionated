package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title   string   `json:"title"`
	PollID  uint     `json:"pollId"`
	Votes   []Vote   `json:"votes"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	QuestionID  uint   `json:"questionId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
