package services

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Wait4Ever(wakems int) {
	go func() {
		for {
			time.Sleep(time.Duration(wakems) * time.Millisecond)
			logrus.Debug("I am waiting patiently ...")
		}
	}()
}
