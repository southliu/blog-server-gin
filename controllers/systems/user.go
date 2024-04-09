package controllers

import (
	"blog-gin/cache"
	"blog-gin/controllers"
	"blog-gin/middleware"
	models "blog-gin/models/systems"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct{}

type UserApi struct {
	ID       int         `json:"id"`
	Username string      `json:"username"`
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	IsFrozen int         `json:"isFrozen"`
	Token    interface{} `json:"token"`
}

func (UserController) GetPostForm(c *gin.Context) (models.User, error) {
	var result models.User

	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	avatar := c.DefaultPostForm("avatar", "")
	email := c.DefaultPostForm("email", "")
	phone := c.DefaultPostForm("phone", "")
	nickname := c.DefaultPostForm("nickname", "")

	if username == "" || password == "" {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return result, gin.Error{}
	}

	result = models.User{
		Username: username,
		Password: controllers.EncryptMd5(password),
		Avatar:   avatar,
		Email:    email,
		Phone:    phone,
		Nickname: nickname,
	}
	return result, nil
}

func (UserController) ReturnUserApi(user models.User, token interface{}) UserApi {
	result := UserApi{
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
		IsFrozen: user.IsFrozen,
		Token:    token,
	}

	return result
}

func (UserController) Register(c *gin.Context) {
	user, err := UserController{}.GetPostForm(c)
	if err != nil {
		return
	}

	configPassword := c.Copy().DefaultPostForm("configPassword", "")
	if user.Password != controllers.EncryptMd5(configPassword) {
		controllers.ReturnError(c, 500, "密码和确认密码不相同")
		return
	}

	sameUser, _ := models.User{}.GetUserByUsername(user.Username)
	if sameUser.ID != 0 {
		controllers.ReturnError(c, 500, "该用户名已注册")
		return
	}

	_, err = models.User{}.Register(user)
	if err != nil {
		controllers.ReturnError(c, 500, "注册失败，请联系管理员")
		return
	}

	controllers.ReturnSuccess(c, 200, "注册成功", UserController{}.ReturnUserApi(user, nil))
}

func (UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	user, _ := models.User{}.GetUserByUsername(username)
	if user.ID == 0 {
		controllers.ReturnError(c, 500, "用户不存在")
		return
	}
	if controllers.EncryptMd5(password) != user.Password {
		controllers.ReturnError(c, 500, "用户名或密码不正确")
		return
	}

	claims := middleware.JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-60 * time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "south",
		},
	}
	tokenStr, err := middleware.CreateJwt(claims)

	cacheKey := "login:" + strconv.FormatUint(user.ID, 10)
	cache.Rab.Set(cache.Rctx, cacheKey, tokenStr, 24*time.Hour)

	if err != nil {
		controllers.ReturnError(c, 500, "生成token失败:"+err.Error())
		return
	}

	controllers.ReturnSuccess(c, 200, "登录成功", UserController{}.ReturnUserApi(user, tokenStr))
}
