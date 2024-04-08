package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint64         `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createAt"`             // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}