package entities

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Tokens       []Token
	WalletType   WalletType
	WalletTypeID uint
	GuestID      uint
}

type WalletType struct {
	gorm.Model
	EventID           uint
	StartValidityDate *time.Time `json:"start_validity_date" gorm:"not null" validate:"required"`
	EndValidityDate   *time.Time `json:"end_validity_date" gorm:"not null" validate:"required"`
	Name              string     `json:"name" gorm:"not null" validate:"required"`
	MaxAmount         int        `json:"max_amount" gorm:"not null" validate:"required"`
	OnlineReload      bool       `json:"online_reload"`
	WalletPricings    []WalletPricing
}

type WalletPricing struct {
	gorm.Model
	Type         string  `json:"type" gorm:"not null" validate:"required"`
	Quantity     int     `json:"quantity" gorm:"not null" validate:"required"`
	UnitPrice    float32 `json:"unit_price" gorm:"not null" validate:"required"`
	WalletTypeID uint
}

type WalletTransaction struct {
	gorm.Model
	CashierID       uint
	WalletPricingID uint
	WalletPricing   WalletPricing
	Quantity        int    `json:"quantity" gorm:"not null" validate:"required"`
	Status          string `json:"type" gorm:"not null" validate:"required"`
}
