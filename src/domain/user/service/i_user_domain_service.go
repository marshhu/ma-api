package service

import "github.com/marshhu/ma-api/src/domain/user/do"

type IUserDomainService interface {
	Get(username string, password string) *do.UserDo
	Add(userName string, password string, gender int64) (id int64, err error)
	IsUserNameValid(userName string) bool
	IsNickNameValid(nickName string) bool
	IsExist(id int) bool
}
