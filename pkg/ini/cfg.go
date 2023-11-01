package ini

import (
	"github.com/joho/godotenv"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"strings"
)

func configInit() {
	env := app.Getenv("MYPASS_ENV", "development")

	env = strings.ToLower(env)
	if env != "" {
		_ = godotenv.Load(".env." + env)
	}
	_ = godotenv.Load()

	app.Cfg = &app.Config{
		Env:             app.ParseEnv(env),
		Host:            app.Getenv("MYPASS_HOST", "0.0.0.0"),
		Port:            app.Getenv("MYPASS_PORT", "7277"),
		SecretKey:       app.Getenv("MYPASS_SECRET_KEY", "super-unsafe-secret-key"),
		LogLevel:        app.ParseLogLevel(app.Getenv("MYPASS_LOGLEVEL", "info")),
		DbConnectionUri: app.Getenv("MYPASS_DB_CONNECTION_URI", "file::memory:?cache=shared"),
	}
}
