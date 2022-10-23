package entities

type User struct {
	Default
	Username          string              `json:"username" validate:"required,min=3,max=32"`
	Email             string              `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password          string              `json:"password" gorm:"not null" validate:"required"`
	Firstname         string              `json:"firstname" gorm:"not null" validate:"required"`
	Lastname          string              `json:"lastname" gorm:"not null" validate:"required"`
	Phone             string              `json:"phone"`
	Events            []Event             `gorm:"many2many:user_events"`
	TicketTransaction []TicketTransaction `gorm:"foreignKey:CashierID"`
	WalletTransaction []WalletTransaction `gorm:"foreignKey:CashierID"`
	TokenTransaction  []TokenTransaction  `gorm:"foreignKey:SellerID"`
}
