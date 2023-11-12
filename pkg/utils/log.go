package utils

import (
	"github.com/sirupsen/logrus"
)

func ParseLogLevel(lvl string) logrus.Level {
	parsed, err := logrus.ParseLevel(lvl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic":   "Logging configuration",
			"event":   "Using default",
			"level":   lvl,
			"default": "info",
		}).Warn("Failed parsing Logger level.")
		return logrus.InfoLevel
	}
	return parsed
}
