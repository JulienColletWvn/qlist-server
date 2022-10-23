package utils

import (
	"fmt"
	"qlist/db/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	env := GetEnvVariable
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Europe/Brussels",
		"localhost", 5432, env("POSTGRES_USER"), env("POSTGRES_PASSWORD"), env("POSTGRES_DB"))

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(
		&entities.User{},
		&entities.Event{},
		&entities.EventImage{},
		&entities.EventProduct{},
		&entities.GuestGroup{},
		&entities.Guest{},
		&entities.Wallet{},
		&entities.WalletType{},
		&entities.WalletPricing{},
		&entities.WalletTransaction{},
		&entities.Token{},
		&entities.TokenTransaction{},
		&entities.Ticket{},
		&entities.TicketType{},
		&entities.TicketTransaction{},
		&entities.Language{},
		&entities.LocalisedTextContent{},
		&entities.LocalisedTextTranslation{},
	)

	return nil
}
