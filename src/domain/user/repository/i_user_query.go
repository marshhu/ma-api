package repository

import "github.com/marshhu/ma-api/src/domain/user/do"

//读仓储
type IUserQuery interface {
	GetById(id int64) *do.UserDo
	Get(username string, password string) *do.UserDo
	GetPage(limit int, offset int) (total int64, list []do.UserDo)
	GetByName(userName string) *do.UserDo
	IsUserNameValid(userName string) bool
	IsNickNameValid(nickName string) bool
	IsExist(id int) bool
}
