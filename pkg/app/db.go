package app

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func init() {
	tmpDB, err := NewDB(Cfg.DbConnectionUri)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	DB = tmpDB
}
