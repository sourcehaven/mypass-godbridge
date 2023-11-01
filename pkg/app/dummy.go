package app

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/db"
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
	"log"
)

func DummyDbInit() {
	orm, err := db.ConnectToMemSQLite()
	if err != nil {
		log.Fatal(err)
	}

	err = orm.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	orm.Create(&models.User{
		Username:  "dummy",
		Email:     "dummy@mail.com",
		Password:  "foobar",
		Firstname: "foo",
		Lastname:  "bar"})
}
