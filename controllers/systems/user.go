package controllers

import (
	"blog-gin/controllers"
	models "blog-gin/models/systems"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UserController) GetPostForm(c *gin.Context) (models.User, error) {
	var res models.User

	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	avatar := c.DefaultPostForm("avatar", "")
	email := c.DefaultPostForm("email", "")
	phone := c.DefaultPostForm("phone", "")
	nickname := c.DefaultPostForm("nickname", "")

	if username == "" || password == "" {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return res, gin.Error{}
	}

	res = models.User{
		Username: username,
		Password: controllers.EncryptMd5(password),
		Avatar:   avatar,
		Email:    email,
		Phone:    phone,
		Nickname: nickname,
	}
	return res, nil
}

// func (UserController) ReturnUserApi(user models.User) UserApi {
// 	var user UserApi

// 	return user
// }

func (UserController) Register(c *gin.Context) {
	user, err := UserController{}.GetPostForm(c)
	if err != nil {
		return
	}

	configPassword := c.Copy().DefaultPostForm("configPassword", "")
	if user.Password != configPassword {
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

	controllers.ReturnSuccess(c, 200, "注册成功", user)
}

func (UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	user, _ := models.User{}.GetUserByUsername(username)
	if user.ID == 0 {
		controllers.ReturnError(c, 500, "用户不存在")
		return
	}
	if password != controllers.EncryptMd5(user.Password) {
		controllers.ReturnError(c, 500, "用户名或密码不正确")
		return
	}

	// TODO: redis操作

	controllers.ReturnSuccess(c, 200, "登录成功", user)
}
