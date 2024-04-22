package main

import (
	"blog-gin/dao"
	"blog-gin/middleware"
	models "blog-gin/models/systems"
	"blog-gin/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := dao.Db.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.Role{},
	)
	if err != nil {
		println(err.Error())
		middleware.Error(map[string]interface{}{"auto migrate db error": err.Error()})
	}

	r := router.Router()
	r.Run(":9999")
}
