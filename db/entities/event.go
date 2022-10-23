package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name          LocalisedTextContent `json:"name" gorm:"polymorphic:Content" validate:"required"`
	Description   string               `json:"description" gorm:"not null" validate:"required"`
	StartDate     *time.Time           `json:"start_date" gorm:"not null" validate:"required"`
	EndDate       *time.Time           `json:"end_date" gorm:"not null" validate:"required"`
	Location      string               `json:"location" gorm:"not null" validate:"required"`
	FreeWifi      bool                 `json:"free_wifi"`
	Public        bool                 `json:"public" gorm:"not null" validate:"required"`
	TicketsAmount int                  `json:"tickets_amount"`
	Images        []EventImage
	Products      []EventProduct
	Cashiers      []User  `gorm:"many2many:events_cashiers"`
	Sellers       []User  `gorm:"many2many:events_sellers"`
	Guests        []Guest `gorm:"many2many:events_guests"`
	TicketTypes   []TicketType
	WalletType    []WalletType
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
