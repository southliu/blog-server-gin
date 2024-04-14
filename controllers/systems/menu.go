package controllers

import (
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
