package dao

import (
	"blog-gin/config"
	"blog-gin/middleware"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MysqlDb), &gorm.Config{})

	if err != nil {
		middleware.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		middleware.Error(map[string]interface{}{"database error": Db.Error})
	}

	// 设置连接池
	sqlDB, err := Db.DB()
	if err != nil {
		middleware.Error(map[string]interface{}{"database pool error": err.Error()})
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
