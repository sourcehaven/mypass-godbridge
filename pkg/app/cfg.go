package app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
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

var Cfg *Config
