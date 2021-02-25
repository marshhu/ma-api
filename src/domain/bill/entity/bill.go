package entity

import "time"

//账单实体
type Bill struct {
	Id          int64     `gorm:"primary_key;column:id"`
	Type        int       `gorm:"column:type"`
	Amount      float64   `gorm:"column:amount"`
	Description string    `gorm:"column:description"`
	BillDate    time.Time `gorm:"column:bill_date"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	CreateBy    string    `gorm:"column:create_by"`
}

func (e *Bill) TableName() string {
	return "bills"
}
