package services

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"time"
)

func (ctx *Context) Wait4Ever(wakems int) {
	if ctx.Config.Env == app.Development {
		go func() {
			for {
				time.Sleep(time.Duration(wakems) * time.Millisecond)
				ctx.Logger.Println("I am waiting patiently ...")
			}
		}()
	}
}
