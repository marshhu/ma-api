package repository

import (
	"time"

	"github.com/marshhu/ma-api/src/domain/user/entity"
	"github.com/marshhu/ma-frame/orm"
)

//用户写仓储实现
type UserRepository struct {
}

func (r *UserRepository) Add(userName string, password string, gender int64) (id int64, err error) {
	user := entity.User{Name: userName, Password: password, Gender: gender, CreateTime: time.Now(), UpdateTime: time.Now()}
	db := orm.NewQuery()
	err = db.Create(&user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *UserRepository) UpdateNickName(id int64, nickName string) error {
	db := orm.NewQuery()
	return db.Model(&entity.User{}).Where("id = ?", id).Update("nick_name", nickName)
}

func (r *UserRepository) UpdateAvatar(id int64, url string) error {
	db := orm.NewQuery()
	return db.Model(&entity.User{}).Where("id = ?", id).Update("avatar", url)
}

func (r *UserRepository) Delete(id int64) error {
	db := orm.NewQuery()
	return db.Delete(&entity.User{}, id)
}
