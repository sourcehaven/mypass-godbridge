package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host      string
	Port      string
	SecretKey string
	Env       string
}

var Cfg *Config

func init() {
	env := getEnvOrPanic("MYPASS_ENV")

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env." + env + " file!")
	}

	Cfg = &Config{
		Host:      getEnvOrPanic("MYPASS_HOST"),
		Port:      getEnvOrPanic("MYPASS_PORT"),
		SecretKey: getEnvOrPanic("MYPASS_SECRET_KEY"),
		Env:       env,
	}
}

func getEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(key + " is not found in environment variables!")
	}
	return val
}
