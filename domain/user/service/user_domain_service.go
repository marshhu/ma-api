package service

import (
	"github.com/marshhu/ma-api/domain/user/do"
	"github.com/marshhu/ma-api/domain/user/repository"
)

type UserDomainService struct {
	UserQuery repository.IUserQuery `inject:"UserQuery"`
}

func (d *UserDomainService) Get(username string, password string) *do.UserDo {
	return d.UserQuery.Get(username, password)
}
