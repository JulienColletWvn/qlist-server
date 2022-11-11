package entities

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	ContactID uint
	Contact   Contact
	Wallets   []Wallet `json:"wallets"`
	Tickets   []Ticket `json:"tickets"`
}

type GuestGroup struct {
	gorm.Model
	Guests []Guest
	Name   string `json:"name" gorm:"not null" validate:"required"`
	Color  string `json:"color"`
}
