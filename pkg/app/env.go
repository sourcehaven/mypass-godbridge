package app

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Getenv(key string, dft string) string {
	val := os.Getenv(key)
	if val == "" {
		logrus.WithFields(logrus.Fields{
			"topic":   "Configuration warning",
			"event":   "Using fallback",
			"key":     key,
			"default": dft,
		}).Warn("Key not found.")
		return dft
	}
	return val
}

func parseEnv(env string) Appenv {
	switch strings.ToLower(env) {
	case "devel", "development":
		return Development
	case "local":
		return Local
	case "stage", "staging":
		return Staging
	case "test", "testing":
		return Testing
	case "prod", "production":
		return Production
	default:
		// use vanilla logrus as Logger is not guaranteed to be configured at this point
		logrus.WithFields(logrus.Fields{
			"topic":   "Application environment configuration",
			"event":   "Using default",
			"level":   env,
			"default": "local",
		}).Warn("Failed parsing environment config.")
		return Local
	}
}
