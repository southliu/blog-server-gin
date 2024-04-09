package main

import (
	"blog-gin/config"
	"blog-gin/dao"
	"blog-gin/middleware"
	models "blog-gin/models/systems"
	"blog-gin/router"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := dao.Db.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.Permission{},
		&models.Role{},
	)
	if err != nil {
		println(err.Error())
		middleware.Error(map[string]interface{}{"auto migrate db error": err.Error()})
	}

	a, err := gormadapter.NewAdapter("mysql", config.MysqlDb, true)
	if err != nil {
		println(err.Error())
		middleware.Error(map[string]interface{}{"db connect error": err.Error()})
	}

	e, err := casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
		middleware.Error(map[string]interface{}{"casbin init error": err.Error()})
	}

	e.LoadPolicy()
	e.Enforce("alice", "data1", "read")
	e.SavePolicy()

	r := router.Router()
	r.Run(":9999")
}
