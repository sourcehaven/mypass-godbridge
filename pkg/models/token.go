package models

import (
	"gorm.io/gorm"
)

type PWToken struct {
	gorm.Model
	ID     uint
	Secret string `gorm:"unique"`
	Salt   string
}

type PPToken struct {
	gorm.Model
	ID     uint
	Secret string `gorm:"unique"`
	Salt   string
}
