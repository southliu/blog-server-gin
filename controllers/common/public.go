package controllers

import (
	"blog-gin/cache"
	"blog-gin/controllers"
	userControllers "blog-gin/controllers/systems"
	"blog-gin/middleware"
	"blog-gin/models/global"
	models "blog-gin/models/systems"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PublicController struct{}

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

	_, err = userModel.Create(&user)
	if err != nil {
		controllers.ReturnError(c, 500, "注册失败，请联系管理员")
		return
	}

	controllers.ReturnSuccess(c, 200, "注册成功", new(userControllers.UserController).ReturnUserApi(user, nil, nil))
}

func (*PublicController) Login(c *gin.Context) {
	var userModel models.User
	var userInfo models.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return
	}

	if userInfo.Username == "" || userInfo.Password == "" {
		controllers.ReturnError(c, 500, userInfo)
		return
	}

	user, _ := userModel.GetUserByUsername(userInfo.Username)
	if user.ID == 0 {
		controllers.ReturnError(c, 500, "用户不存在")
		return
	}
	if controllers.EncryptMd5(userInfo.Password) != user.Password {
		controllers.ReturnError(c, 500, "用户名或密码不正确")
		return
	}

	// 获取改用户下面的角色
	roles, err := new(models.User).GetRoleByUsername(userInfo.Username)
	if err != nil {
		controllers.ReturnError(c, 500, "获取角色信息失败")
		return
	}

	var roleTokens []uint64
	var permissions []string
	for _, role := range roles {
		rolePermissions, _ := new(models.Role).GetPermissionById(role.ID)
		roleTokens = append(roleTokens, role.ID)

		for _, permission := range rolePermissions {
			if permission != "" {
				permissions = append(permissions, permission)
			}
		}
	}

	claims := middleware.JwtClaims{
		Username: userInfo.Username,
		UserId:   userInfo.ID,
		Roles:    roleTokens,
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

	controllers.ReturnSuccess(
		c,
		200,
		"登录成功",
		new(userControllers.UserController).ReturnUserApi(user, tokenStr, permissions),
	)
}

// 刷新权限
func (*PublicController) RefreshPermission(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	authHeaderLen := len(authHeader)
	tokenStr := authHeader[7:authHeaderLen]

	tokenInfo, err := middleware.GetJwtInfo(authHeader)
	if err != nil {
		controllers.ReturnError(c, 401, "当前权限已失效，请重新登录")
		return
	}
	username := tokenInfo.(*middleware.JwtClaims).Username

	// 获取用户
	user, err := new(models.User).GetUserByUsername(username)
	if err != nil {
		controllers.ReturnError(c, 500, "获取用户信息失败")
		return
	}

	// 获取改用户下面的角色
	roles, err := new(models.User).GetRoleByUsername(username)
	if err != nil {
		controllers.ReturnError(c, 500, "获取角色信息失败")
		return
	}

	var permissions []string
	for _, role := range roles {
		rolePermissions, _ := new(models.Role).GetPermissionById(role.ID)
		for _, permission := range rolePermissions {
			if permission != "" {
				permissions = append(permissions, permission)
			}
		}
	}

	result := new(userControllers.UserController).ReturnUserApi(user, tokenStr, permissions)

	controllers.ReturnSuccess(
		c,
		200,
		"登录成功",
		result,
	)
}

func (*PublicController) Init(c *gin.Context) {
	menuData := []*models.Menu{
		{
			GVA_MODEL: global.GVA_MODEL{
				ID: 1,
			},
			Label:      "系统管理",
			LabelEn:    "system",
			Type:       0,
			SortNum:    1,
			Permission: "/system",
		},
		{
			GVA_MODEL: global.GVA_MODEL{
				ID: 2,
			},
			PId:        1,
			Label:      "菜单管理",
			LabelEn:    "menu",
			Type:       1,
			SortNum:    1,
			Route:      "/system/menu",
			Permission: "/system/menu",
		},
		{PId: 2, Label: "菜单管理-查看", Type: 2, SortNum: 1, Permission: "/system/menu/index"},
		{PId: 2, Label: "菜单管理-新增", Type: 2, SortNum: 2, Permission: "/system/menu/create"},
		{PId: 2, Label: "菜单管理-编辑", Type: 2, SortNum: 3, Permission: "/system/menu/update"},
		{PId: 2, Label: "菜单管理-删除", Type: 2, SortNum: 4, Permission: "/system/menu/delete"},
	}
	menus, err := new(models.Menu).BatchCreate(menuData)
	if err != nil {
		controllers.ReturnError(c, 500, err.Error())
		return
	}

	// Casbin权限
	middleware.Casbin.AddPolicy("/system/menu/index", "/systems/menu/list", "GET")
	middleware.Casbin.AddPolicy("/system/menu/index", "/systems/menu/detail", "GET")
	middleware.Casbin.AddPolicy("/system/menu/create", "/systems/menu", "POST")
	middleware.Casbin.AddPolicy("/system/menu/update", "/systems/menu", "PUT")
	middleware.Casbin.AddPolicy("/system/menu/delete", "/systems/menu", "DELETE")

	// 初始化角色
	role1 := models.Role{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		Name:      "管理员角色",
	}
	role2 := models.Role{
		GVA_MODEL: global.GVA_MODEL{ID: 2},
		Name:      "游客角色",
	}
	findRole1, _ := new(models.Role).GetRoleById(1)
	if findRole1.ID != 0 {
		controllers.ReturnError(c, 500, "角色存在")
		return
	}
	findRole2, _ := new(models.Role).GetRoleById(2)
	if findRole2.ID != 0 {
		controllers.ReturnError(c, 500, "角色存在")
		return
	}

	for _, value := range menus {
		if value.ID != 0 {
			role1.Menus = append(role1.Menus, *value)
		}
	}
	newRole1, _ := new(models.Role).Create(&role1)
	newRole2, _ := new(models.Role).Create(&role2)

	// 初始化用户
	user1 := models.User{
		Username: "admin",
		Password: controllers.EncryptMd5("admin123456"),
		Nickname: "管理员",
	}
	user1.Roles = append(user1.Roles, *newRole1)
	user2 := models.User{
		Username: "south",
		Password: controllers.EncryptMd5("south123456"),
		Nickname: "游客",
	}
	user2.Roles = append(user2.Roles, *newRole2)
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

	new(models.User).Create(&user1)
	_, err = new(models.User).Create(&user2)
	if err != nil {
		controllers.ReturnError(c, 500, err.Error())
		return
	}

	controllers.ReturnSuccess(c, 200, "success", "初始化成功")
}
