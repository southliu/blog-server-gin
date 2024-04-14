package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type Menu struct {
	global.GVA_MODEL
	Name       string `json:"name" gorm:"type:varchar(50);comment:名称;not null" binding:"required"`
	Type       int    `json:"type" gorm:"comment:类型:0—目录 1—菜单 2—按钮;not null"`
	Route      string `json:"route" gorm:"type:varchar(100);comment:路由"`
	Icon       string `json:"icon" gorm:"type:varchar(100);comment:图标"`
	SortNum    int    `json:"sortNum" gorm:"type:tinyint;comment:排序"`
	Enable     int    `json:"enable" gorm:"type:tinyint;default:1;comment:是否启用 0禁用 1启用"`
	Permission string `json:"permission" gorm:"type:varchar(200);comment:权限"`
	PId        uint64 `json:"pid" gorm:"default:0;comment:父级ID"`
	global.GVA_DATE_MODEL

	Roles []Role `json:"roles" gorm:"many2many:sys_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MenuPageSearch struct {
	global.PAGE_MODEL
}

func (Menu) TableName() string {
	return "sys_menus"
}

func (*Menu) Create(menu Menu) (Menu, error) {
	err := dao.Db.Create(&menu).Error
	return menu, err
}

func (*Menu) BatchCreate(menus []Menu) ([]Menu, error) {
	err := dao.Db.Create(&menus).Error
	return menus, err
}

func (*Menu) GetMenuPage(search MenuPageSearch) ([]Menu, error) {
	var list []Menu
	page := search.PAGE_MODEL.Page
	pageSize := search.PAGE_MODEL.PageSize
	offset := (page - 1) * pageSize
	err := dao.Db.Offset(offset).Limit(pageSize).Find(&list).Error
	return list, err
}
