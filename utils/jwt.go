package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecretkey"

func GenerateJWT(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : email,
		"userId" : userId,
		"exp" : time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signin method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token!")
	}
	
	if !parsedToken.Valid {
		return errors.New("Token not valid!")
	}

	// claims, ok := parsedToken.Claims.(*jwt.MapClaims)
	
	// if !ok {
	// 	return errors.New("Invalid token claims!")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}