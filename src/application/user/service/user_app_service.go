package service

import (
	"errors"

	"github.com/marshhu/ma-api/src/application/user/ao"
	userDomainSrv "github.com/marshhu/ma-api/src/domain/user/service"
	"github.com/marshhu/ma-frame/utils"
)

type UserAppService struct {
	UserDomainService userDomainSrv.IUserDomainService `inject:"UserDomainService"`
}

func (s *UserAppService) GetById(userId int64) *ao.UserAo {
	userDo := s.UserDomainService.GetById(userId)
	var userAo ao.UserAo
	utils.MapTo(userDo, &userAo)
	return &userAo
}

func (s *UserAppService) Get(username string, password string) *ao.UserAo {
	userDo := s.UserDomainService.Get(username, password)
	var userAo ao.UserAo
	utils.MapTo(userDo, &userAo)
	return &userAo
}

func (s *UserAppService) Add(input ao.RegisterAo) (id int64, err error) {
	if !s.UserDomainService.IsUserNameValid(input.UserName) {
		return 0, errors.New("用户名已存在")
	}
	return s.UserDomainService.Add(input.UserName, input.Password, 0)
}
