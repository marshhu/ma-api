package service

import (
	"github.com/marshhu/ma-api/src/application/user/ao"
)

type IUserAppService interface {
	Get(username string, password string) *ao.UserAo
	Add(input ao.RegisterAo) (id int64, err error)
}
