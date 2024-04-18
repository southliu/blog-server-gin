package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type Menu struct {
	global.GVA_MODEL
	Label      string `json:"label" gorm:"type:varchar(50);comment:名称;not null" binding:"required"`
	LabelEn    string `json:"labelEn" gorm:"type:varchar(50);comment:名称;"`
	Type       int    `json:"type" gorm:"comment:类型:0—目录 1—菜单 2—按钮;not null"`
	Route      string `json:"route" gorm:"type:varchar(100);comment:路由"`
	Icon       string `json:"icon" gorm:"type:varchar(100);comment:图标"`
	SortNum    int    `json:"sortNum" gorm:"type:tinyint;comment:排序"`
	Enable     int    `json:"enable" gorm:"type:tinyint;default:1;comment:是否启用 0禁用 1启用"`
	Permission string `json:"permission" gorm:"type:varchar(200);comment:权限"`
	PId        uint64 `json:"pid" gorm:"default:0;comment:父级ID"`
	Key        string `json:"key" gorm:"-"`
	Children   []Menu `json:"children" gorm:"-"`
	global.GVA_DATE_MODEL

	Roles []Role `json:"roles" gorm:"many2many:sys_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MenuPageSearch struct {
	global.PAGE_MODEL
}

func (Menu) TableName() string {
	return "sys_menus"
}

func (*Menu) Create(menu *Menu) (*Menu, error) {
	err := dao.Db.Create(&menu).Error
	return menu, err
}

func (*Menu) BatchCreate(menus []*Menu) ([]*Menu, error) {
	err := dao.Db.Create(&menus).Error
	return menus, err
}

func (*Menu) GetMenuPage(search MenuPageSearch) ([]Menu, error) {
	var list []Menu
	page := search.PAGE_MODEL.Page
	pageSize := search.PAGE_MODEL.PageSize
	offset := (page - 1) * pageSize
	err := dao.Db.Unscoped().Offset(offset).Limit(pageSize).Find(&list).Error
	return list, err
}

func (*Menu) GetMenuList(roleIds []uint64) ([]Menu, error) {
	var roles []Role
	var menus []Menu
	err := dao.Db.Unscoped().Model(&Role{}).Preload("Menus").Where("id IN ?", roleIds).Find(&roles).Error
	if err == nil {
		for _, role := range roles {
			menus = append(menus, role.Menus...)
		}
	}
	return menus, err
}

func (*Menu) GetMenuById(id uint64) (Menu, error) {
	var menu Menu
	err := dao.Db.Unscoped().Where("id = ?", id).First(&menu).Error
	return menu, err
}

func (*Menu) Update(id uint64, menu *Menu) (*Menu, error) {
	err := dao.Db.Model(Menu{}).Where("id = ?", id).Updates(&menu).Error
	return menu, err
}
