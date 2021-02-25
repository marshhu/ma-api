package ioc

import (
	"github.com/facebookgo/inject"
	billAppSrv "github.com/marshhu/ma-api/src/application/bill/service"
	userAppSrv "github.com/marshhu/ma-api/src/application/user/service"
)

// DIContainer
var DIContainer Container

// Container
type Container struct {
	UserAppService userAppSrv.IUserAppService `inject:"UserAppService"`
	BillAppService billAppSrv.IBillAppService `inject:"BillAppService"`
}

// ioc
func InitIoc() {
	var g inject.Graph
	handleErr(g.Provide(&inject.Object{Value: &DIContainer}))
	initUserIoc(&g)
	initBillIoc(&g)
	handleErr(g.Populate())
}

//错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
