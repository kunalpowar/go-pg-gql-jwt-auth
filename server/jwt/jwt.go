package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("jwt: error in generating key: %v", err)
	}

	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}

	return "", err
}
