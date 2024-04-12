package models

import (
	"blog-gin/models/global"
)

type UserRole struct {
	global.GVA_MODEL
	UserId uint64 `gorm:"index;comment:关联用户ID"`
	RoleId uint64 `gorm:"index;comment:关联角色ID"`
}

func (*UserRole) TableName() string {
	return "sys_user_roles"
}
