package entity

import "time"

//用户实体
type User struct {
	Id         int64     `gorm:"primary_key;column:id"`
	Name       string    `gorm:"column:name"`
	Password   string    `gorm:"column:password"`
	NickName   string    `gorm:"column:nick_name"`
	Gender     int64     `gorm:"column:gender"`
	Avatar     string    `gorm:"column:avatar"`
	Tel        string    `gorm:"column:tel"`
	QQ         string    `gorm:"column:qq"`
	WeChat     string    `gorm:"column:we_chat"`
	Email      string    `gorm:"column:email"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (e *User) TableName() string {
	return "users"
}
