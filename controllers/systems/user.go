package controllers

import (
	"blog-gin/controllers"
	models "blog-gin/models/systems"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

/** 获取post请求参数 */
func (*UserController) GetPostForm(c *gin.Context) (models.User, error) {
	var result models.User

	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	avatar := c.DefaultPostForm("avatar", "")
	email := c.DefaultPostForm("email", "")
	phone := c.DefaultPostForm("phone", "")
	nickname := c.DefaultPostForm("nickname", "")

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

/** 返回API数据 */
func (UserController) ReturnUserApi(user models.User, token interface{}) models.UserApi {
	result := models.UserApi{
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
