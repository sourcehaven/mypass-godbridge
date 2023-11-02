package ini

import (
	"github.com/joho/godotenv"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/security/jwt"
	"strings"
	"time"
)

func configInit() {
	env := app.Getenv("MYPASS_ENV", "development")

	env = strings.ToLower(env)
	if env != "" {
		_ = godotenv.Load(".env." + env)
	}
	_ = godotenv.Load()

	const host = "0.0.0.0"
	const port = "7277"
	const secret = "super-unsafe-secret-key"
	const algo = "HS256"
	const level = "info"
	const dburi = "file::memory:?cache=shared"

	app.Cfg = &app.Config{
		Env:               app.ParseEnv(env),
		Host:              app.Getenv("MYPASS_HOST", host),
		Port:              app.Getenv("MYPASS_PORT", port),
		SecretKey:         app.Getenv("MYPASS_SECRET_KEY", secret),
		JwtAccessKey:      app.Getenv("MYPASS_JWT_ACCESS_KEY", app.Getenv("MYPASS_SECRET_KEY", secret)),
		JwtRefreshKey:     app.Getenv("MYPASS_JWT_REFRESH_KEY", app.Getenv("MYPASS_SECRET_KEY", secret)),
		JwtSigningMethod:  jwt.ParseJwtSigningMethod(app.Getenv("MYPASS_JWT_SIGNING_METHOD", algo)),
		JwtAccessExpires:  10 * time.Minute,
		JwtRefreshExpires: 240 * time.Hour,
		LogLevel:          app.ParseLogLevel(app.Getenv("MYPASS_LOGLEVEL", level)),
		DbConnectionUri:   app.Getenv("MYPASS_DB_CONNECTION_URI", dburi),
	}
}
