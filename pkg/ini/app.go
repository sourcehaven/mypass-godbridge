package ini

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/services"
	"os/exec"
)

func InitApp() {
	if app.Cfg.Env == app.Development {
		dummyDbInit()

		// swag init --generalInfo=./cmd/mypass/main.go
		prog := "swag"
		cmd := exec.Command(prog, "init", "--generalInfo=./cmd/mypass/main.go")
		stdout, err := cmd.Output()
		logrus.WithFields(logrus.Fields{
			"out": fmt.Sprintf("%s", stdout),
			"err": err,
		}).Info()
	}
	// initialize services
	services.Wait4Ever(60000) // dummy service logs every 10 minute
}
