package ioc

import (
	"github.com/facebookgo/inject"
	billAppSrv "github.com/marshhu/ma-api/src/application/bill/service"
	billRepo "github.com/marshhu/ma-api/src/domain/bill/repository"
	billDomainSrv "github.com/marshhu/ma-api/src/domain/bill/service"
)

//Bill
func initBillIoc(g *inject.Graph) {
	handleErr(g.Provide(
		&inject.Object{Value: &billAppSrv.BillAppService{}, Name: "BillAppService"},
		&inject.Object{Value: &billDomainSrv.BillDomainService{}, Name: "BillDomainService"},
		&inject.Object{Value: &billRepo.BillQuery{}, Name: "BillQuery"},
		&inject.Object{Value: &billRepo.BillRepository{}, Name: "BillRepository"},
	))
}
