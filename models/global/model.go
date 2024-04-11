package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID uint64 `gorm:"primarykey" json:"id"` // 主键ID
}

type GVA_Date_MODEL struct {
	CreatedAt time.Time      `json:"createAt"`        // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`        // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"` // 删除时间
}

type GVA_Middle_Date_MODEL struct {
	CreatedAt time.Time `json:"createAt"` // 创建时间
	UpdatedAt time.Time `json:"updateAt"` // 更新时间
}
