package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ParseJwtSigningMethod(method string) jwt.SigningMethod {
	switch strings.ToLower(method) {
	case "hs256":
		return jwt.SigningMethodHS256
	case "hs384":
		return jwt.SigningMethodHS384
	case "hs512":
		return jwt.SigningMethodHS512
	case "es256":
		return jwt.SigningMethodES256
	case "es384":
		return jwt.SigningMethodES384
	case "es512":
		return jwt.SigningMethodES512
	case "rs256":
		return jwt.SigningMethodRS256
	case "rs384":
		return jwt.SigningMethodRS384
	case "rs512":
		return jwt.SigningMethodRS512
	case "ps256":
		return jwt.SigningMethodPS256
	case "ps384":
		return jwt.SigningMethodPS384
	case "ps512":
		return jwt.SigningMethodPS512
	case "eddsa":
		return jwt.SigningMethodEdDSA
	case "none", "":
		logrus.WithFields(logrus.Fields{
			"topic": "Empty signature algorithm",
			"sign":  method,
		}).Warn("Moving on without signature algorithm. This is usually a bad idea. :(")
		return jwt.SigningMethodNone
	default:
		logrus.WithFields(logrus.Fields{
			"topic":   "Signature algorithm",
			"sign":    method,
			"default": "hs256",
		}).Warn("Selecting default signing algorithm.")
		return jwt.SigningMethodHS256
	}
}

type TokenOptions struct {
	Payload           Payload
	Fresh             Fresh
	Refresh           Refresh
	JwtAccessExpires  time.Duration
	JwtRefreshExpires time.Duration
	JwtSigningMethod  jwt.SigningMethod
	JwtAccessKey      string
	JwtRefreshKey     string
	AllowExpired      bool
}

func CreateAccessToken(identity string, opts *TokenOptions) (tokenString string, err error) {
	claims := &Claim{
		Fresh:   opts.Fresh,
		Payload: opts.Payload,
		Refresh: false,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   identity,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(opts.JwtAccessExpires).UTC()),
		},
	}
	token := jwt.NewWithClaims(opts.JwtSigningMethod, claims)
	base64Key := []byte(opts.JwtAccessKey)
	tokenString, err = token.SignedString(base64Key)
	return
}

func CreateRefreshToken(identity string, opts *TokenOptions) (tokenString string, err error) {
	claims := &Claim{
		Payload:          opts.Payload,
		Refresh:          true,
		RegisteredClaims: jwt.RegisteredClaims{Subject: identity},
	}
	token := jwt.NewWithClaims(opts.JwtSigningMethod, claims)
	base64Key := []byte(opts.JwtRefreshKey)
	tokenString, err = token.SignedString(base64Key)
	return
}

var (
	ErrTokenExpired   = errors.New("token expired")
	ErrTokenMalformed = errors.New("malformed claims")
)

func commonChecks(token jwt.Token, claims *Claim, err *error, checkExpiry bool) {
	if *err != nil {
		return
	}
	claim, ok := token.Claims.(*Claim)
	*claims = *claim
	if !ok {
		*err = ErrTokenMalformed
		return
	}
	if checkExpiry && claims.ExpiresAt != nil && claims.ExpiresAt.Unix() < time.Now().Unix() {
		*err = ErrTokenExpired
		return
	}
}

func ValidateAccessToken(signedToken string, opts *TokenOptions) (claims *Claim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(opts.JwtAccessKey), nil
		},
	)
	claims = &Claim{}
	commonChecks(*token, claims, &err, !opts.AllowExpired)
	if err != nil {
		claims = nil
		return
	}
	if claims.Refresh {
		err = errors.New("expected access token")
		return
	}
	return
}

func ValidateFreshAccessToken(signedToken string, opts *TokenOptions) (claims *Claim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(opts.JwtAccessKey), nil
		},
	)
	claims = &Claim{}
	commonChecks(*token, claims, &err, !opts.AllowExpired)
	if err != nil {
		claims = nil
		return
	}
	if claims.Refresh {
		err = errors.New("expected access token")
		return
	}
	if !claims.Fresh {
		err = errors.New("fresh access token required")
		return
	}
	return
}

func ValidateRefreshToken(signedToken string, opts *TokenOptions) (claims *Claim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(opts.JwtRefreshKey), nil
		},
	)
	claims = &Claim{}
	commonChecks(*token, claims, &err, !opts.AllowExpired)
	if err != nil {
		claims = nil
		return
	}
	if !claims.Refresh {
		err = errors.New("expected refresh token")
		return
	}
	return
}
