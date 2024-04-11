package controllers

import (
	"blog-gin/cache"
	"blog-gin/controllers"
	userControllers "blog-gin/controllers/systems"
	"blog-gin/middleware"
	models "blog-gin/models/systems"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PublicController struct{}

func (*PublicController) Init(c *gin.Context) {
	// 初始化角色
	role1 := models.Role{
		Name: "管理员角色",
	}
	role2 := models.Role{
		Name: "游客角色",
	}
	// findRole1, _ := new(models.Role).GetRoleById(1)
	// if findRole1.ID != 0 {
	// 	controllers.ReturnError(c, 500, "角色存在")
	// 	return
	// }
	// findRole2, _ := new(models.Role).GetRoleById(2)
	// if findRole2.ID != 0 {
	// 	controllers.ReturnError(c, 500, "角色存在")
	// 	return
	// }
	newRole1, _ := new(models.Role).Create(role1)
	newRole2, _ := new(models.Role).Create(role2)

	// 初始化用户
	user1 := models.User{
		Username: "admin",
		Password: "admin666",
		Nickname: "管理员",
	}
	user1.Roles = append(user1.Roles, &newRole1)
	user2 := models.User{
		Username: "south",
		Password: "south666",
		Nickname: "游客",
	}
	user2.Roles = append(user2.Roles, &newRole2)
	finUser1, _ := new(models.User).GetUserByUsername(user1.Username)
	if finUser1.ID != 0 {
		controllers.ReturnError(c, 500, "用户存在")
		return
	}

	finUser2, _ := new(models.User).GetUserByUsername(user2.Username)
	if finUser2.ID != 0 {
		controllers.ReturnError(c, 500, "用户存在")
		return
	}

	_, err := new(models.User).Create(user1)
	if err != nil {
		controllers.ReturnError(c, 500, err.Error())
		return
	}
	new(models.User).Create(user2)

	controllers.ReturnSuccess(c, 200, "success", "初始化成功")
}

func (*PublicController) Register(c *gin.Context) {
	var userModel models.User
	user, err := new(userControllers.UserController).GetPostForm(c)
	if err != nil {
		return
	}

	if user.Username == "" || user.Password == "" {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return
	}

	configPassword := c.Copy().DefaultPostForm("configPassword", "")
	if user.Password != controllers.EncryptMd5(configPassword) {
		controllers.ReturnError(c, 500, "密码和确认密码不相同")
		return
	}

	sameUser, _ := userModel.GetUserByUsername(user.Username)
	if sameUser.ID != 0 {
		controllers.ReturnError(c, 500, "该用户名已注册")
		return
	}

	_, err = userModel.Create(user)
	if err != nil {
		controllers.ReturnError(c, 500, "注册失败，请联系管理员")
		return
	}

	controllers.ReturnSuccess(c, 200, "注册成功", new(userControllers.UserController).ReturnUserApi(user, nil))
}

func (*PublicController) Login(c *gin.Context) {
	var userModel models.User
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return
	}

	user, _ := userModel.GetUserByUsername(username)
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

	controllers.ReturnSuccess(c, 200, "登录成功", new(userControllers.UserController).ReturnUserApi(user, tokenStr))
}
