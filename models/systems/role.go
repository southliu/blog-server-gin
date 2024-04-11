package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type Role struct {
	global.GVA_MODEL
	Name   string `json:"name" gorm:"column:name;type:varchar(50);comment:角色名;not null;" binding:"required"`
	UserId uint64 `json:"-" gorm:"column:user_id;type:int;comment:关联用户ID"`
	global.GVA_Date_MODEL

	Permissions []*Permission `json:"permissions" gorm:"many2many:sys_role_permissions;"`
	Users       []*User       `gorm:"many2many:sys_user_roles;"`
}

func (*Role) TableName() string {
	return "sys_roles"
}

func (*Role) Create(role Role) (Role, error) {
	err := dao.Db.Create(&role).Error
	return role, err
}

func (*Role) GetRoleById(id uint64) (Role, error) {
	var role Role
	err := dao.Db.Where("id = ?", id).First(&role).Error
	return role, err
}
