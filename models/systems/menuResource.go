package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type MenuResource struct {
	global.GVA_MODEL
	MenuId uint64 `json:"menuId" gorm:"comment:菜单ID;not null" binding:"required"`
	Method string `json:"method" gorm:"varchar(20);comment:请求方法"`
	Path   string `json:"path" gorm:"varchar(255);comment:请求路径"`
	global.GVA_DATE_MODEL
}

func (MenuResource) TableName() string {
	return "sys_menu_resource"
}

func (*MenuResource) Create(menuResource *MenuResource) (*MenuResource, error) {
	err := dao.Db.Create(&menuResource).Error
	return menuResource, err
}

func (*MenuResource) BatchCreate(menuResources []*MenuResource) ([]*MenuResource, error) {
	err := dao.Db.Create(&menuResources).Error
	return menuResources, err
}
