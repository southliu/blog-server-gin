package models

import "blog-gin/models/global"

type Role struct {
	global.GVA_MODEL
	Name string ` json:"name" gorm:"column:name;type:varchar(50);comment:角色名;not null;many2many:sys_permissions;" binding:"required"`
	global.GVA_Date_MODEL
}

func (Role) TableName() string {
	return "sys_roles"
}
