package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := make(jwt.MapClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, "", errors.New("invalid token")
	}
	return int64(id), claims["username"].(string), nil
}
