package utils

import (
	"errors"
	"time"

	"github.com/aldyN25/myproject/app/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID string) (string, error) {
	configs := config.GetInstance()
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(configs.Jwtconfig.Secret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return result, nil
}

func GenerateAccessToken(userID string) (string, error) {
	return GenerateToken(userID)
}

func GenerateRefreshToken(userID string) (string, error) {
	return GenerateToken(userID)
}
