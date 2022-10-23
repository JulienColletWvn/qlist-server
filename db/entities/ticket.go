package entities

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Uuid         string `json:"uuid" validate:"required"`
	Transactions []TicketTransaction
	TicketType   TicketType
	TicketTypeID uint
	GuestID      uint
}

type TicketType struct {
	gorm.Model
	EventID           uint
	Name              string     `json:"name" gorm:"not null" validate:"required"`
	StartValidityDate *time.Time `json:"start_validity_date" gorm:"not null" validate:"required"`
	EndValidityDate   *time.Time `json:"end_validity_date" gorm:"not null" validate:"required"`
	Quantity          int        `json:"quantity"`
	Unlimited         bool       `json:"unlimited"`
}

type TicketTransaction struct {
	gorm.Model
	TicketID  uint
	CashierID uint
	Quantity  int    `json:"quantity" gorm:"not null" validate:"required"`
	Status    string `json:"type" gorm:"not null" validate:"required"`
}
