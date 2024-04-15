package controllers

import (
	"blog-gin/controllers"
	"blog-gin/middleware"
	"blog-gin/models/global"
	models "blog-gin/models/systems"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

func (*MenuController) GetMenuPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultPostForm("page", "10"))

	search := models.MenuPageSearch{
		PAGE_MODEL: global.PAGE_MODEL{
			Page:     page,
			PageSize: pageSize,
		},
	}

	new(models.Menu).GetMenuPage(search)
}

func (*MenuController) BuildTree(menus []models.Menu, pID uint64) []models.Menu {
	var tree []models.Menu

	for _, menu := range menus {
		if menu.PId == pID && menu.Type < 2 {
			children := new(MenuController).BuildTree(menus, menu.ID)
			if len(children) > 0 {
				menu.Children = children
			}
			menu.Key = menu.Route
			tree = append(tree, menu)
		}
	}

	return tree
}

func (*MenuController) GetMenuList(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenInfo := middleware.GetJwtInfo(authHeader)
	roles := tokenInfo.(*middleware.JwtClaims).Roles

	menus, err := new(models.Menu).GetMenuList(roles)
	if err != nil {
		controllers.ReturnError(c, 500, "获取菜单列表失败"+err.Error())
		return
	}
	controllers.ReturnSuccess(c, 200, "success", new(MenuController).BuildTree(menus, 0))
}
