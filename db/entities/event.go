package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Content       []LocalisedTextContent `json:"content" gorm:"polymorphic:Content" validate:"required"`
	StartDate     *time.Time             `json:"start_date" gorm:"not null" validate:"required"`
	EndDate       *time.Time             `json:"end_date" gorm:"not null" validate:"required"`
	Location      string                 `json:"location" gorm:"not null" validate:"required"`
	FreeWifi      bool                   `json:"free_wifi"`
	Public        bool                   `json:"public" gorm:"not null"`
	TicketsAmount int                    `json:"tickets_amount"`
	Status        string                 `json:"status"`
	Images        []EventImage           `json:"images"`
	Products      []EventProduct         `json:"products"`
	Cashiers      []User                 `json:"cashiers" gorm:"many2many:events_cashiers"`
	Sellers       []User                 `json:"sellers" gorm:"many2many:events_sellers"`
	Guests        []Guest                `json:"guests" gorm:"many2many:events_guests"`
	TicketTypes   []TicketType           `json:"tickets"`
	WalletTypes   []WalletType           `json:"wallets"`
}

type EventImage struct {
	gorm.Model
	EventID uint
	Url     string `json:"url" gorm:"not null" validate:"required"`
}

type EventProduct struct {
	gorm.Model
	EventID    uint
	Name       string `json:"name" gorm:"polymorphic:Content" validate:"required"`
	TokenPrice int    `json:"token_price" gorm:"not null" validate:"required"`
}
