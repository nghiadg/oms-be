package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GetExpAccessToken() int64 {
	return time.Now().Add(time.Hour).Unix()
}

func GetExpRefreshToken() int64 {
	return time.Now().Add(time.Hour * 8).Unix()
}

type ClaimsAccessToken struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

type ClaimsRefreshToken struct {
	UserId int `json:"user_id"`
}

func GenerateAccessToken(claims ClaimsAccessToken) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.UserId,
		"email":   claims.Email,
		"exp":     GetExpAccessToken(),
	})

	return accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_KEY")))
}

func GenerateRefreshToken(claims ClaimsRefreshToken) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.UserId,
		"exp":     GetExpRefreshToken(),
	})

	return refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_KEY")))
}
