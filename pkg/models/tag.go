package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	UserID      uint
	Name        string
	Description string
}
