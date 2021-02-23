package service

import (
	"github.com/marshhu/ma-api/application/user/ao"
)

type IUserAppService interface {
	Get(username string, password string) *ao.UserAo
}
