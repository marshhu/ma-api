package repository

import (
	"github.com/marshhu/ma-api/src/domain/user/do"
	"github.com/marshhu/ma-api/src/domain/user/entity"
	"github.com/marshhu/ma-frame/orm"
	"github.com/marshhu/ma-frame/utils"
)

//用户读仓储实现
type UserQuery struct {
}

func (q *UserQuery) GetById(id int64) *do.UserDo {
	var user entity.User
	db := orm.NewQuery()
	err := db.Model(&entity.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil
	}
	var result do.UserDo
	utils.MapTo(&user, &result)
	return &result
}

func (q *UserQuery) Get(username string, password string) *do.UserDo {
	var user entity.User
	db := orm.NewQuery()
	err := db.Model(&entity.User{}).Where("name = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		return nil
	}
	var result do.UserDo
	utils.MapTo(&user, &result)
	return &result
}

func (q *UserQuery) GetPage(limit int, offset int) (total int64, list []do.UserDo) {
	var users []entity.User
	db := orm.NewQuery()
	err := db.Model(&entity.User{}).Order("create_time desc").Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return 0, nil
	}
	utils.MapTo(&users, &list)
	db.Model(&entity.User{}).Count(&total)
	return total, list
}

func (q *UserQuery) GetByName(userName string) *do.UserDo {
	var user entity.User
	db := orm.NewQuery()
	err := db.Model(&entity.User{}).Where("name = ?", userName).First(&user).Error
	if err != nil {
		return nil
	}
	var result do.UserDo
	utils.MapTo(&user, &result)
	return &result
}

func (q *UserQuery) IsUserNameValid(userName string) bool {
	db := orm.NewQuery()
	var count int64
	db.Model(&entity.User{}).Where("name = ?", userName).Count(&count)
	return count <= 0
}

func (q *UserQuery) IsNickNameValid(nickName string) bool {
	db := orm.NewQuery()
	var count int64
	db.Model(&entity.User{}).Where("name_name = ?", nickName).Count(&count)
	return count <= 0
}

func (q *UserQuery) IsExist(id int) bool {
	db := orm.NewQuery()
	var count int64
	db.Model(&entity.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}
