package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type CustomClaims struct {
	ID uint
	jwt.RegisteredClaims
}

func GenerateToken(id uint) (tokenstring string, err error) {
	claims := CustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, _ = token.SignedString(secretKey)
	return
}

func ReadToken(tokenstring string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenstring, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	return
}

