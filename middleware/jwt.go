package middleware

import (
	"blog-gin/config"
	"blog-gin/controllers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserId   uint64   `json:"userId"`
	Username string   `json:"username"`
	Roles    []uint64 `json:"roles"`
	jwt.RegisteredClaims
}

func JwtAuth(c *gin.Context) {
	// 获取请求URI
	url := c.Request.URL.Path
	// 获取请求方法
	act := c.Request.Method

	for _, item := range config.Whitelists {
		if url == item.Url && act == item.Act {
			c.Next()
			return
		}
	}

	authHeader := c.Request.Header.Get("Authorization")
	tokenInfo, err := GetJwtInfo(authHeader)
	if err != nil {
		controllers.ReturnError(c, 401, "当前权限已失效，请重新登录")
		c.Abort()
		return
	}

	userId := tokenInfo.(*JwtClaims).UserId
	username := tokenInfo.(*JwtClaims).Username
	roles := tokenInfo.(*JwtClaims).Roles
	c.Set("userId", userId)
	c.Set("username", username)
	c.Set("roles", roles)

	c.Next()
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
