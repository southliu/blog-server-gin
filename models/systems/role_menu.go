package models

type RoleMenu struct {
	RoleId uint64 `gorm:"index;comment:关联角色ID"`
	MenuId uint64 `gorm:"index;comment:关联菜单ID"`
}

func (*RoleMenu) TableName() string {
	return "sys_role_menus"
}
