package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID uint64 `json:"id" gorm:"primarykey"` // 主键ID
}

type GVA_DATE_MODEL struct {
	CreatedAt time.Time      `json:"createAt"`        // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`        // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"` // 删除时间
}

type PAGE_MODEL struct {
	Page     int `json:"page"`     // 当前页数
	PageSize int `json:"pageSize"` // 当前分页总条数
}
