package entities

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	Email        string `json:"email" gorm:"not null" validate:"required,email"`
	Password     string `json:"password" gorm:"not null" validate:"required"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Phone        string `json:"phone"`
	Wallets      []Wallet
	Tickets      []Ticket
	GuestGroupID uint
}

type GuestGroup struct {
	gorm.Model
	Guests []Guest
	Name   string `json:"name" gorm:"not null" validate:"required"`
	Color  string `json:"color"`
}
