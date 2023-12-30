package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Username string
	jwt.StandardClaims
}

var secretKey = os.Getenv("SECRET_KEY")

func GenerateAccessToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		panic(err.Error())
	}

	return tokenString
}

func ParseToken(signedToken string) *TokenClaims {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		panic(err.Error())
	}

	claims, ok := token.Claims.(*TokenClaims)
	fmt.Println(">>>", claims, ok)

	if !ok || err != nil {
		panic("Couldn't parse token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		panic("Token expired")
	}

	return claims
}
