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

	isMenuSelectStr := c.DefaultQuery("isMenuSelect", "false")
	isMenuSelect := isMenuSelectStr == "true"

	menus, err := new(models.Menu).GetMenuList(roles)
	if err != nil {
		controllers.ReturnError(c, 500, "获取菜单列表失败")
		return
	}
	result := new(MenuController).BuildTree(menus, 0, isAll)

	if isMenuSelect {
		result = []models.Menu{
			{
				GVA_MODEL: global.GVA_MODEL{
					ID: 0,
				},
				Label:    "顶级菜单",
				SortNum:  0,
				Enable:   1,
				Children: result,
			},
		}
	}

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

func (*MenuController) Create(c *gin.Context) {
	var menu models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		controllers.ReturnError(c, 500, "请输入正确信息")
		return
	}

	// 获取角色数据
	_roles, _ := c.Get("roles")
	roles := _roles.([]uint64)

	for _, role := range roles {
		roleData, err := new(models.Role).GetRoleById(role)

		if err != nil {
			controllers.ReturnError(c, 500, "新增菜单失败，无法获取对应角色数据")
			return
		}

		menu.Roles = append(menu.Roles, roleData)
	}

	result, err := new(models.Menu).Create(&menu)
	if err != nil {
		controllers.ReturnError(c, 500, "新增菜单失败")
		return
	}

	controllers.ReturnSuccess(c, 200, "新增成功", result)
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
