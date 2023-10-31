package app

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Config struct {
	Host      string
	Port      string
	SecretKey string
	Env       string
}

var Cfg *Config

func init() {
	env := os.Getenv("MYPASS_ENV")
	env = strings.ToLower(env)
	if env != "" {
		_ = godotenv.Load(".env." + env)
	}
	_ = godotenv.Load()

	Cfg = &Config{
		Host:      os.Getenv("MYPASS_HOST"),
		Port:      os.Getenv("MYPASS_PORT"),
		SecretKey: os.Getenv("MYPASS_SECRET_KEY"),
		Env:       env,
	}
}
