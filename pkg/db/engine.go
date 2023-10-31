package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToMemSQLite() (*gorm.DB, error) {
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
