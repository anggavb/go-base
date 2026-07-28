package jwthandler

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}

func New(userId, jwtIssuer string, expHour int) *Claims {
	return &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId, // implement RFC 7519 standard for subject claim
			Issuer:    jwtIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}

func (c *Claims) GenerateJWT(jwtSecret string) (string, error) {
	if jwtSecret == "" {
		return "", ErrMissingJwtSecret
	}
	uToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return uToken.SignedString([]byte(jwtSecret))
}

func (c *Claims) VerifyJWT(token string, jwtSecret string) error {
	if jwtSecret == "" {
		return ErrMissingJwtSecret
	}

	jwtToken, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if !jwtToken.Valid {
		return jwt.ErrTokenInvalidClaims
	}

	issuer, err := jwtToken.Claims.GetIssuer()
	if err != nil {
		return err
	}

	if issuer != c.Issuer {
		return jwt.ErrTokenInvalidClaims
	}

	return nil
}
