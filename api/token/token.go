package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const signingKey = "SecretKeyForJWT"

func GenerateToken(userId uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(signingKey))
}

func Validate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
