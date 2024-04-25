package router

import (
	commonControllers "blog-gin/controllers/common"
	sysControllers "blog-gin/controllers/systems"
	"blog-gin/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(middleware.LoggerToFile()))
	r.Use(middleware.Recover)
	r.Use(middleware.JwtAuth)
	r.Use(middleware.Authorize)

	r.GET("/init", new(commonControllers.PublicController).Init)
	r.POST("/login", new(commonControllers.PublicController).Login)
	r.POST("/register", new(commonControllers.PublicController).Register)
	r.GET("/refresh-permissions", new(commonControllers.PublicController).RefreshPermission)

	systems := r.Group("/systems")
	{
		menu := systems.Group("/menu")
		{
			menu.GET("/page", new(sysControllers.MenuController).GetMenuPage)
			menu.GET("/list", new(sysControllers.MenuController).GetMenuList)
			menu.GET("/detail", new(sysControllers.MenuController).GetMenuById)
			menu.POST("", new(sysControllers.MenuController).Create)
			menu.PUT("/:id", new(sysControllers.MenuController).Update)
			menu.DELETE("/:id", new(sysControllers.MenuController).Delete)
		}
		role := systems.Group("/role")
		{
			role.POST("/create", new(sysControllers.RoleController).Create)
		}
	}
	return r
}
