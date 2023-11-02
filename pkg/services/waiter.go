package services

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"time"
)

func Wait4Ever(wakems int) {
	if app.Cfg.Env == app.Development {
		go func() {
			for {
				time.Sleep(time.Duration(wakems) * time.Millisecond)
				app.Logger.Println("I am waiting patiently ...")
			}
		}()
	}
}
