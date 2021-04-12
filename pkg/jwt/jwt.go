package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secret")

type AccessToken struct {
	Jwt string
	Exp time.Time
}

func GenerateToken(userId string) (*AccessToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = userId

	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["exp"] = exp

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		Jwt: tokenString,
		Exp: time.Unix(exp, 0),
	}, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)

		return userId, nil
	} else {
		return "", err
	}
}
