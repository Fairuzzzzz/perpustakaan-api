package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID int64, role, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	},
	)

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, "", "", err
	}

	if !token.Valid {
		return 0, "", "", err
	}

	userID, okUserID := claims["userID"].(float64)
	role, okRole := claims["role"].(string)
	username, okUsername := claims["username"].(string)
	if !okUserID || !okRole || !okUsername {
		return 0, "", "", errors.New("invalid token claims")
	}
	return int64(userID), role, username, nil
}
