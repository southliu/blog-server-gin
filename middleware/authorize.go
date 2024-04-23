package middleware

import (
	"blog-gin/config"
	"blog-gin/controllers"
	"fmt"
	"log"
	"strconv"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

var (
	Casbin *casbin.Enforcer
)

func init() {
	db, err := gormadapter.NewAdapter("mysql", config.MysqlDb, true)
	if err != nil {
		println(err.Error())
		Error(map[string]interface{}{"db connect error": err.Error()})
	}

	Casbin, err = casbin.NewEnforcer("config/rbac_model.conf", db)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
		Error(map[string]interface{}{"casbin init error": err.Error()})
	}

	Casbin.LoadPolicy()
	Casbin.Enforce("alice", "data1", "read")
	Casbin.SavePolicy()
}

func Authorize(c *gin.Context) {
	// 获取请求URI
	url := c.Request.URL.Path
	// 获取请求方法
	act := c.Request.Method

	for _, item := range config.Whitelists {
		if url == item.Url && act == item.Act {
			c.Next()
			return
		}
	}

	authHeader := c.Request.Header.Get("Authorization")
	tokenInfo, err := GetJwtInfo(authHeader)
	if err != nil {
		controllers.ReturnError(c, 401, "当前权限已失效，请重新登录")
		c.Abort()
		return
	}

	roles := tokenInfo.(*JwtClaims).Roles
	roleIdStr := strconv.FormatUint(roles[0], 10)

	if ok, _ := Casbin.Enforce(roleIdStr, url, act); ok {
		fmt.Println("权限通过" + act + url)
		c.Next()
	} else {
		fmt.Println("权限未通过" + act + url)
		controllers.ReturnError(c, 500, "当前用户无权访问，请重新登录尝试！")
		c.Abort()
	}
}
