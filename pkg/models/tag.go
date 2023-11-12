package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	UserID      uint   `gorm:"uniqueIndex:uid_name"`
	Name        string `gorm:"uniqueIndex:uid_name"`
	Description string
}
