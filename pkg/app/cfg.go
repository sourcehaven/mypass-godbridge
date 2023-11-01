package app

import (
	"github.com/sirupsen/logrus"
)

type Appenv uint32

const (
	Development = iota
	Local
	Staging
	Testing
	Production
)

type Config struct {
	Host            string
	Port            string
	SecretKey       string
	Env             Appenv
	LogLevel        logrus.Level
	DbConnectionUri string
}

var Cfg *Config
