package service

import "github.com/marshhu/ma-api/domain/user/do"

type IUserDomainService interface {
	Get(username string, password string) *do.UserDo
}
