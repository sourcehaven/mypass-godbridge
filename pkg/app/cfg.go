package app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	myjwt "github.com/sourcehaven/mypass-godbridge/pkg/security/jwt"
	"strings"
	"time"
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
	Host              string            // specifies host that the application will run on
	Port              string            // specifies which port the app should run on
	SecretKey         string            // universal secret key used for signing
	JwtAccessKey      string            // secret key for signing access tokens
	JwtRefreshKey     string            // secret key for signing refresh tokens
	JwtAccessExpires  time.Duration     // expiration time of access jwt (should be relatively short-lived)
	JwtRefreshExpires time.Duration     // expiration time of refresh jwt
	JwtSigningMethod  jwt.SigningMethod // method used when signing jwt keys
	Env               Appenv            // current environment eg.: devel, prod, etc...
	LogLevel          logrus.Level      // logging level (trace, debug, info, warn, error)
	DbConnectionUri   string            // specifies the db connection string
}

func NewConfig() *Config {
	env := Getenv("MYPASS_ENV", "development")

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

	return &Config{
		Env:               ParseEnv(env),
		Host:              Getenv("MYPASS_HOST", host),
		Port:              Getenv("MYPASS_PORT", port),
		SecretKey:         Getenv("MYPASS_SECRET_KEY", secret),
		JwtAccessKey:      Getenv("MYPASS_JWT_ACCESS_KEY", Getenv("MYPASS_SECRET_KEY", secret)),
		JwtRefreshKey:     Getenv("MYPASS_JWT_REFRESH_KEY", Getenv("MYPASS_SECRET_KEY", secret)),
		JwtSigningMethod:  myjwt.ParseJwtSigningMethod(Getenv("MYPASS_JWT_SIGNING_METHOD", algo)),
		JwtAccessExpires:  10 * time.Minute,
		JwtRefreshExpires: 240 * time.Hour,
		LogLevel:          ParseLogLevel(Getenv("MYPASS_LOGLEVEL", level)),
		DbConnectionUri:   Getenv("MYPASS_DB_CONNECTION_URI", dburi),
	}
}
