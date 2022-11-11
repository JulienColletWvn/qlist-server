package entities

import "gorm.io/gorm"

type LocalisedTextContent struct {
	gorm.Model
	Lang        string `json:"lang" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Content     string `json:"content" validate:"required"`
	ContentID   int
	ContentType string
}
