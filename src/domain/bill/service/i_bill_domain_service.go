package service

import (
	"github.com/marshhu/ma-api/src/domain/bill/do"
	"github.com/marshhu/ma-api/src/domain/bill/entity"
)

type IBillDomainService interface {
	Add(bill entity.Bill) (id int64, err error)
	Delete(id int64) error
	Update(bill entity.Bill) error
	GetById(id int64) *do.BillDo
	GetByUser(userName string, limit int, offset int) (total int64, list []do.BillDo)
}
