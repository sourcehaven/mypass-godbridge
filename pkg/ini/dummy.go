package ini

import "C"
import (
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
)

func dummyDbInit() {
	err := app.DB.AutoMigrate(&models.User{})
	if err != nil {
		logrus.Fatal(err)
	}
	app.DB.Create(&models.User{
		Username:  "dummy",
		Email:     "dummy@mail.com",
		Password:  "foobar",
		Firstname: "foo",
		Lastname:  "bar",
	})
}
