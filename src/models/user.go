package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	HashedPassword string `json:"hashedPassword"`
	Username       string `json:"username"`
}
