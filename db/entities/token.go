package entities

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	WalletID     uint
	Uuid         string `json:"uuid" validate:"required"`
	Transactions []TokenTransaction
}

type TokenTransaction struct {
	gorm.Model
	Amount   int    `json:"amount" gorm:"not null" validate:"required"`
	Status   string `json:"status" gorm:"not null" validate:"required"`
	SellerID uint
	TokenID  uint
}
