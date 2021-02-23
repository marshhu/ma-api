package repository

import "github.com/marshhu/ma-api/domain/user/do"

//读仓储
type IUserQuery interface {
	Get(username string, password string) *do.UserDo
	GetPage(limit int, offset int) (total int64, list []do.UserDo)
	GetByName(userName string) *do.UserDo
	IsUserNameValid(userName string) bool
	IsNickNameValid(nickName string) bool
	IsExist(id int) bool
}
