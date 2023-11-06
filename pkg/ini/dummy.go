package ini

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/db"
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
)

func (ctx *Context) dummyDbInit() {
	if ctx.Config.Env == app.Development {
		orm, err := db.CreateEngine(ctx.Config.DbConnectionUri)
		if err != nil {
			ctx.Logger.Fatal(err)
		}

		err = orm.AutoMigrate(&models.User{})
		if err != nil {
			ctx.Logger.Fatal(err)
		}
		orm.Create(&models.User{
			Username:  "dummy",
			Email:     "dummy@mail.com",
			Password:  "foobar",
			Firstname: "foo",
			Lastname:  "bar",
		})
	}
}
