package ioc

import (
	"github.com/facebookgo/inject"
	userAppSrv "github.com/marshhu/ma-api/src/application/user/service"
)

// DIContainer
var DIContainer Container

// Container
type Container struct {
	UserAppService userAppSrv.IUserAppService `inject:"UserAppService"`
}

// ioc
func InitIoc() {
	var g inject.Graph
	handleErr(g.Provide(&inject.Object{Value: &DIContainer}))
	initUserIoc(&g)
	handleErr(g.Populate())
}

//错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
