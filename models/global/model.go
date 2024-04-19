package global

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LocalTime sql.NullTime

type GVA_MODEL struct {
	ID uint64 `json:"id" gorm:"primarykey"` // 主键ID
}

type GVA_DATE_MODEL struct {
	CreatedAt LocalTime      `json:"createdAt"`       // 创建时间
	UpdatedAt LocalTime      `json:"updatedAt"`       // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"` // 删除时间
}

type PAGE_MODEL struct {
	Page     int `json:"page"`     // 当前页数
	PageSize int `json:"pageSize"` // 当前分页总条数
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t.Time)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (n *LocalTime) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

func (n LocalTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
