package entities

import "gorm.io/gorm"

type LocalisedTextContent struct {
	gorm.Model
	Translations []LocalisedTextTranslation
	ContentID    int
	ContentType  string
}

type LocalisedTextTranslation struct {
	gorm.Model
	LocalisedTextContentID uint
	Content                string   `json:"content" validate:"required"`
	Language               Language `gorm:"polymorphic:Language" validate:"required"`
}

type Language struct {
	gorm.Model
	Name         string
	LanguageID   int
	LanguageType string
}
