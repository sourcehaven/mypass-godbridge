package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateEngine(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
