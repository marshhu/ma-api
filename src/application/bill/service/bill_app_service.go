package service

import (
	"time"

	"github.com/marshhu/ma-frame/utils"

	"github.com/marshhu/ma-api/src/application/bill/ao"
	"github.com/marshhu/ma-api/src/domain/bill/entity"
	billDomainSrv "github.com/marshhu/ma-api/src/domain/bill/service"
)

type BillAppService struct {
	BillDomainService billDomainSrv.IBillDomainService `inject:"BillDomainService"`
}

func (s *BillAppService) Add(bill ao.AddBillAo) (id int64, err error) {
	billEnt := entity.Bill{
		Amount:      bill.Amount,
		Type:        bill.Type,
		Description: bill.Description,
		BillDate:    utils.ToTime(bill.BillDate, utils.DATE_LAYOUT),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateBy:    bill.CreateBy,
	}
	return s.BillDomainService.Add(billEnt)
}

func (s *BillAppService) GetById(id int64) *ao.BillAo {
	bill := s.BillDomainService.GetById(id)
	var result ao.BillAo
	utils.MapTo(&bill, &result)
	return &result
}

func (s *BillAppService) GetByUser(userName string, limit int, offset int) (total int64, list []ao.BillAo) {
	total, bills := s.BillDomainService.GetByUser(userName, limit, offset)
	utils.MapTo(&bills, &list)
	return
}
