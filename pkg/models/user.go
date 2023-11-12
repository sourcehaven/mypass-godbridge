package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Firstname string
	Lastname  string
	Active    bool `gorm:"default:false"`
}
