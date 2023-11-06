package deps

import (
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
)

type AppContext struct {
	Config *app.Config
	Logger *logrus.Logger
}
