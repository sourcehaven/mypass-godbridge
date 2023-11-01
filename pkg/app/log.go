package app

import (
	"github.com/sirupsen/logrus"
)

func ParseLogLevel(lvl string) logrus.Level {
	parsed, err := logrus.ParseLevel(lvl)
	if err != nil {
		// use vanilla logrus as Logger is not guaranteed to be configured at this point
		Logger.WithFields(logrus.Fields{
			"topic":   "Logging configuration",
			"event":   "Using default",
			"level":   lvl,
			"default": "info",
		}).Warn("Failed parsing Logger level.")
		return logrus.InfoLevel
	}
	return parsed
}

var Logger = logrus.New()
