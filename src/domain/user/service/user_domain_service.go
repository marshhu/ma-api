package service

import (
	"github.com/marshhu/ma-api/src/domain/user/do"
	"github.com/marshhu/ma-api/src/domain/user/repository"
)

type UserDomainService struct {
	UserQuery      repository.IUserQuery      `inject:"UserQuery"`
	UserRepository repository.IUserRepository `inject:"UserRepository"`
}

func (s *UserDomainService) GetById(userId int64) *do.UserDo {
	return s.UserQuery.GetById(userId)
}

func (s *UserDomainService) Get(username string, password string) *do.UserDo {
	return s.UserQuery.Get(username, password)
}

func (s *UserDomainService) Add(userName string, password string, gender int64) (id int64, err error) {
	return s.UserRepository.Add(userName, password, gender)
}

func (s *UserDomainService) IsUserNameValid(userName string) bool {
	return s.UserQuery.IsUserNameValid(userName)
}

func (s *UserDomainService) IsNickNameValid(nickName string) bool {
	return s.UserQuery.IsNickNameValid(nickName)
}

func (s *UserDomainService) IsExist(id int) bool {
	return s.UserQuery.IsExist(id)
}
