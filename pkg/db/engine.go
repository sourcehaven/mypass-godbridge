package db

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateEngine() (*gorm.DB, error) {
	dsn := app.Cfg.DbConnectionUri
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
