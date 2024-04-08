package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwt(claims jwt.Claims) (string, error) {
	mySigningKey := []byte("IAmSouth")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)

	return tokenStr, err
}

func GetJwtInfo(tokenString string) interface{} {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("IAmSouth"), nil
	})
	if err != nil {
		return nil
	} else if claims, ok := token.Claims.(*JwtClaims); ok {
		return claims
	}

	return nil
}
