package entities

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Email     string `json:"email" gorm:"not null" validate:"required,email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	UserID    uint
}
