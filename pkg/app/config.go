package app

import "os"

type Config struct{}

func (cfg *Config) Host() string {
	return os.Getenv("MYPASS_HOST")
}

func (cfg *Config) Port() string {
	return os.Getenv("MYPASS_PORT")
}

func (cfg *Config) SecretKey() string {
	return os.Getenv("MYPASS_SECRET_KEY")
}

func (cfg *Config) Env() string {
	return os.Getenv("MYPASS_ENV")
}
