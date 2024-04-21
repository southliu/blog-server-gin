package global

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Time time.Time

type GVA_MODEL struct {
	ID uint64 `json:"id" gorm:"primarykey"` // 主键ID
}

type GVA_DATE_MODEL struct {
	CreatedAt Time           `json:"createdAt"`       // 创建时间
	UpdatedAt Time           `json:"updatedAt"`       // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"` // 删除时间
}

type PAGE_MODEL struct {
	Page     int `json:"page"`     // 当前页数
	PageSize int `json:"pageSize"` // 当前分页总条数
}

// 1. 为 Xtime 重写 UnmarshalJSON 方法，在此方法中实现自定义格式的转换；
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*t = Time(time.Unix(int64(num), 0))
	return
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))), nil
	// return ([]byte)(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// 3. 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).Unix() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

// 4. 为 Time 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
