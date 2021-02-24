package ioc

import (
	"github.com/facebookgo/inject"
	userAppSrv "github.com/marshhu/ma-api/src/application/user/service"
	userRepo "github.com/marshhu/ma-api/src/domain/user/repository"
	userDomainSrv "github.com/marshhu/ma-api/src/domain/user/service"
)

//User
func initUserIoc(g *inject.Graph) {
	handleErr(g.Provide(
		&inject.Object{Value: &userAppSrv.UserAppService{}, Name: "UserAppService"},
		&inject.Object{Value: &userDomainSrv.UserDomainService{}, Name: "UserDomainService"},
		&inject.Object{Value: &userRepo.UserQuery{}, Name: "UserQuery"},
		&inject.Object{Value: &userRepo.UserRepository{}, Name: "UserRepository"},
	))
}
