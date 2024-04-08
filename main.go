package main

import (
	"blog-gin/dao"
	"blog-gin/middleware"
	models "blog-gin/models/systems"
	"blog-gin/router"
)

func main() {
	err := dao.Db.AutoMigrate(&models.User{})
	if err != nil {
		println(err)
		middleware.Error(map[string]interface{}{"auto migrate db error": err.Error()})
	}

	r := router.Router()
	r.Run(":9999")
}
