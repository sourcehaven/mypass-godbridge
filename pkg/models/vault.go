package models

import "gorm.io/gorm"

type Vault struct {
	gorm.Model
	UserID   uint
	Username string
	Password string
	Title    string
	Website  string
	Notes    string
	Folder   string
	Active   bool
	Deleted  bool

	ParentID *uint
	Parent   *Vault `gorm:"foreignkey:ParentID"`
	TagID    *uint
	Tags     []Tag `gorm:"foreignkey:TagID"`
}
