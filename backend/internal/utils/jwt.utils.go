package utils

import (
	"backend/internal/helpers"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user_id int64) (string, error) {
	helpers.LoadEnvironmentFile()

	expStr := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	secret := os.Getenv("ACCESS_TOKEN_SECRET")

	expDuration, err := time.ParseDuration(expStr)

	if err != nil {
		log.Fatal("Error parsing duration for access token")
	}

	claims := jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(expDuration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(secret))
}

func ValidateAccessToken(tokenStr string) (*jwt.Token, error) {
	helpers.LoadEnvironmentFile()
	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, nil
}

func GenerateRefreshToken(user_id int64) (string, error) {
	helpers.LoadEnvironmentFile()

	expStr := os.Getenv("REFRESH_TOKEN_EXPIRATION")
	secret := os.Getenv("REFREESH_TOKEN_SECRET")

	expDuration, err := time.ParseDuration(expStr)

	if err != nil {
		log.Fatal("Error parsing duration for access token")
	}

	claims := jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(expDuration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(secret))
}

func ValidateRefreshToken(tokenStr string) (*jwt.Token, error) {
	helpers.LoadEnvironmentFile()
	secret := os.Getenv("REFRESH_TOKEN_SECRET")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, nil
}
