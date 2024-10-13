package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// Claims for access token
type AccessClaims struct {
	Issuer   string `json:"iss"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// RefreshClaims for refresh token
type RefreshClaims struct {
	Issuer   string `json:"iss"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string) (string, error) {
	claims := AccessClaims{
		Issuer:   "access",
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // 15 minutes expiration
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("jwt_secret_key")))
}

func GenerateRefreshToken(username string) (string, error) {
	claims := RefreshClaims{
		Issuer:   "refresh",
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 7 days expiration
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("jwt_secret_key")))
}
