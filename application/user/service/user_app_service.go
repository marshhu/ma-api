package service

import (
	"github.com/marshhu/ma-api/application/user/ao"
	userDomainSrv "github.com/marshhu/ma-api/domain/user/service"
	"github.com/marshhu/ma-frame/utils"
)

type UserAppService struct {
	UserDomainService userDomainSrv.IUserDomainService `inject:"UserDomainService"`
}

func (a *UserAppService) Get(username string, password string) *ao.UserAo {
	userDo := a.UserDomainService.Get(username, password)
	var userAo ao.UserAo
	utils.MapTo(userDo, &userAo)
	return &userAo
}
