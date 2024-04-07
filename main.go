package main

import (
	"blog-gin/dao"
	models "blog-gin/models/systems"
	"blog-gin/pkg/logger"
	"blog-gin/router"
)

func main() {
	err := dao.Db.AutoMigrate(&models.User{})
	if err != nil {
		println(err)
		logger.Error(map[string]interface{}{"auto migrate db error": err.Error()})
	}

	r := router.Router()
	r.Run(":9999")
}
