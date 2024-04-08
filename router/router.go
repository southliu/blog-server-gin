package router

import (
	controllers "blog-gin/controllers/systems"
	"blog-gin/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(middleware.LoggerToFile()))
	r.Use(middleware.Recover)

	r.POST("/login", controllers.UserController{}.Login)
	r.POST("/register", controllers.UserController{}.Register)

	// user := r.Group("/user")
	// {
	// 	user.POST("/register", controllers.UserController{}.Register)
	// 	user.POST("/login", controllers.UserController{}.Login)
	// }
	return r
}
