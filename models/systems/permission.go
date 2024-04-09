package models

import (
	"blog-gin/models/global"
)

type Permission struct {
	global.GVA_MODEL
	Code        string `json:"code" gorm:"column:code;type:varchar(50);comment:权限代码;not null" binding:"required"`
	Description string `json:"description" gorm:"column:description;type:varchar(100);comment:权限描述;"`
	MenuId      int    `json:"-"`
	global.GVA_Date_MODEL
}

func (Permission) TableName() string {
	return "sys_permissions"
}
