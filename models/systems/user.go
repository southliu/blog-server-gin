package models

import (
	"blog-gin/dao"
	"blog-gin/models/global"
)

type User struct {
	global.GVA_MODEL
	Username string `json:"username" gorm:"type:varchar(100);comment:用户名;not null" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(100);comment:密码;not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(50);comment:昵称"`
	Avatar   string `json:"avatar" gorm:"type:varchar(100);comment:头像"`
	Email    string `json:"email" gorm:"type:varchar(100);comment:邮箱"`
	Phone    string `json:"phone" gorm:"type:varchar(100);comment:电话"`
	IsFrozen int    `json:"isFrozen" gorm:"default:0;comment:是否冻结 0正常 1冻结"`
	global.GVA_DATE_MODEL

	Roles []Role `gorm:"many2many:sys_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserApi struct {
	UserInfo    `json:"userInfo"`
	Permissions []string    `json:"permissions"`
	Token       interface{} `json:"token"`
}

type UserPageSearch struct {
	global.PAGE_MODEL
}

func (User) TableName() string {
	return "sys_users"
}

func (*User) Create(user User) (User, error) {
	err := dao.Db.Create(&user).Error
	return user, err
}

func (*User) BatchCreate(users []User) ([]User, error) {
	err := dao.Db.Create(&users).Error
	return users, err
}

func (*User) GetUserByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Unscoped().Where("username = ?", username).First(&user).Error
	return user, err
}

func (*User) GetUserPage(search UserPageSearch) ([]User, error) {
	var list []User
	page := search.PAGE_MODEL.Page
	pageSize := search.PAGE_MODEL.PageSize
	offset := (page - 1) * pageSize
	err := dao.Db.Offset(offset).Limit(pageSize).Find(&list).Error
	return list, err
}

func (*User) GetRoleByUsername(username string) ([]Role, error) {
	var user User
	err := dao.Db.Unscoped().Preload("Roles").Where("username = ?", username).First(&user).Error
	return user.Roles, err
}
