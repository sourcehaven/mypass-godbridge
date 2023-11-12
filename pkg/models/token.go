package models

import (
	"gorm.io/gorm"
)

// PWToken
// Model for storing password encrypted user token
type PWToken struct {
	gorm.Model
	ID     uint
	Secret string `gorm:"unique"`
	Salt   string
}

// PPToken
// Model for storing passphrase encrypted user token
type PPToken struct {
	gorm.Model
	ID     uint
	Secret string `gorm:"unique"`
	Salt   string
}
