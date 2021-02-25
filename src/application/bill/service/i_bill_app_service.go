package service

import (
	"github.com/marshhu/ma-api/src/application/bill/ao"
)

type IBillAppService interface {
	Add(bill ao.AddBillAo) (id int64, err error)
	GetById(id int64) *ao.BillAo
	GetByUser(userName string, limit int, offset int) (total int64, list []ao.BillAo)
}
