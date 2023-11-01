package ini

import "github.com/sourcehaven/mypass-godbridge/pkg/app"

func loggerInit() {
	app.Logger.SetLevel(app.Cfg.LogLevel)
}
