package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Username string   `json:"username"`
	Roles    []uint64 `json:"roles"`
	jwt.RegisteredClaims
}

func CreateJwt(claims jwt.Claims) (string, error) {
	mySigningKey := []byte("IAmSouth")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)

	return tokenStr, err
}

func GetJwtInfo(tokenString string) (interface{}, error) {
	tokenStringLen := len(tokenString)
	newTokenString := tokenString[7:tokenStringLen]
	token, err := jwt.ParseWithClaims(newTokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("IAmSouth"), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*JwtClaims); ok {
		return claims, nil
	}

	return nil, nil
}
