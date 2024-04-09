package models

import (
	"blog-gin/models/global"
)

type Menu struct {
	global.GVA_MODEL
	Name       string     ` json:"name" gorm:"column:name;type:varchar(50);comment:名称;not null" binding:"required"`
	Type       int        ` json:"type" gorm:"column:type;type:int;comment:类型:0—目录 1—菜单 2—按钮;not null"`
	Route      string     ` json:"route" gorm:"column:route;type:varchar(100);comment:路由"`
	Icon       string     ` json:"icon" gorm:"column:icon;type:varchar(100);comment:图标"`
	SortNum    int        ` json:"sortNum" gorm:"column:sortNum;type:int;comment:排序"`
	Enable     int        ` json:"enable" gorm:"column:enable;type:int;comment:是否启用 0禁用 1启用"`
	Permission Permission `json:"permissions"`
	global.GVA_Date_MODEL
}

func (Menu) TableName() string {
	return "sys_menus"
}
