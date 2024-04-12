package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type Role struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"type:varchar(50);comment:角色名;not null;" binding:"required"`
	global.GVA_Date_MODEL

	Menus []*Menu `gorm:"many2many:sys_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Users []*User `gorm:"many2many:sys_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (*Role) TableName() string {
	return "sys_roles"
}

func (*Role) Create(role Role) (Role, error) {
	err := dao.Db.Create(&role).Error
	return role, err
}

func (*Role) BatchCreate(roles []Role) ([]Role, error) {
	err := dao.Db.Create(&roles).Error
	return roles, err
}

func (*Role) GetRoleById(id uint64) (Role, error) {
	var role Role
	err := dao.Db.Where("id = ?", id).First(&role).Error
	return role, err
}
