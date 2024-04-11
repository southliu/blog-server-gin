package models

import (
	"blog-gin/models/global"
)

type RolePermission struct {
	global.GVA_MODEL
	RoleId       uint64 `gorm:"column:role_id;type:int;comment:关联角色ID"`
	PermissionId uint64 `gorm:"column:permission_id;type:int;comment:关联权限ID"`
	global.GVA_Middle_Date_MODEL
}

func (*RolePermission) TableName() string {
	return "sys_role_permissions"
}
