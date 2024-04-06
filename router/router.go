package router

import (
	"blog-gin/controllers"
	"blog-gin/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	r.POST("/login", controllers.UserController{}.Login)

	// user := r.Group("/user")
	// {
	// 	user.POST("/register", controllers.UserController{}.Register)
	// 	user.POST("/login", controllers.UserController{}.Login)
	// }
	return r
}
