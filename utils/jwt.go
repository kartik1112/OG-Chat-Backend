package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key string

func GenerateJWTToken(email string, userId int) (string, error) {
	key = os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(key))
}

func VerifyJWTToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected JWT signing Method")
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("unexpected JWT signing Method")
	}
	isValid := parsedToken.Valid
	if !isValid {
		return "", errors.New("invalid Token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		if !ok {
			return "", errors.New("invalid Token Claim")
		}

	}
	email := claims["email"].(string)
	return email, nil

}
