package models

import (
	"blog-gin/models/global"
)

type UserRole struct {
	global.GVA_MODEL
	UserId uint64 `gorm:"column:user_id;type:int;comment:关联用户ID"`
	RoleId uint64 `gorm:"column:role_id;type:int;comment:关联角色ID"`
	global.GVA_Middle_Date_MODEL
}

func (*UserRole) TableName() string {
	return "sys_user_roles"
}
