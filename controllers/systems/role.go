package controllers

import (
	"blog-gin/controllers"
	models "blog-gin/models/systems"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (*RoleController) Create(c *gin.Context) {
	var role models.Role
	name := c.DefaultPostForm("name", "")
	role.Name = name
	_, err := role.Create(&role)

	if err != nil {
		controllers.ReturnError(c, 500, "新增失败，请联系管理员")
	}
}
