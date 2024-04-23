package controllers

import (
	"blog-gin/controllers"
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

func (*MenuController) BuildTree(menus []models.Menu, pID uint64, isAll bool) []models.Menu {
	var tree []models.Menu

	for _, menu := range menus {
		if menu.PId == pID && ((!isAll && menu.Type < 2) || isAll) {
			children := new(MenuController).BuildTree(menus, menu.ID, isAll)
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
	_roles, _ := c.Get("roles")
	roles := _roles.([]uint64)

	isAllStr := c.DefaultQuery("isAll", "false")
	isAll := isAllStr == "true"

	menus, err := new(models.Menu).GetMenuList(roles)
	if err != nil {
		controllers.ReturnError(c, 500, "获取菜单列表失败")
		return
	}
	result := new(MenuController).BuildTree(menus, 0, isAll)
	controllers.ReturnSuccess(c, 200, "success", result)
}

func (*MenuController) GetMenuById(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		controllers.ReturnError(c, 500, "请输入ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controllers.ReturnError(c, 500, "请正确ID")
		return
	}

	result, err := new(models.Menu).GetMenuById(id)
	if err != nil {
		controllers.ReturnError(c, 500, "获取菜单数据错误")
		return
	}

	controllers.ReturnSuccess(c, 200, "success", result)
}

func (*MenuController) Update(c *gin.Context) {
	var menu models.Menu
	idStr := c.Param("id")
	if idStr == "" {
		controllers.ReturnError(c, 500, "请输入ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controllers.ReturnError(c, 500, "请正确ID")
		return
	}

	if err := c.ShouldBindJSON(&menu); err != nil {
		controllers.ReturnError(c, 500, "请输入正确信息"+err.Error())
		return
	}

	result, err := new(models.Menu).Update(id, &menu)
	if err != nil {
		controllers.ReturnError(c, 500, "修改菜单失败")
		return
	}
	controllers.ReturnSuccess(c, 200, "修改成功", result)
}

func (*MenuController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		controllers.ReturnError(c, 500, "请输入ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controllers.ReturnError(c, 500, "请正确ID")
		return
	}

	err = new(models.Menu).Delete(id)
	if err != nil {
		controllers.ReturnError(c, 500, "修改菜单失败"+err.Error())
		return
	}
	controllers.ReturnSuccess(c, 200, "删除成功", "")
}
