package service

import (
	"github.com/marshhu/ma-api/src/domain/bill/do"
	"github.com/marshhu/ma-api/src/domain/bill/entity"
	"github.com/marshhu/ma-api/src/domain/bill/repository"
)

type BillDomainService struct {
	BillQuery      repository.IBillQuery      `inject:"BillQuery"`
	BillRepository repository.IBillRepository `inject:"BillRepository"`
}

func (s *BillDomainService) Add(bill entity.Bill) (id int64, err error) {
	return s.BillRepository.Add(bill)
}

func (s *BillDomainService) Delete(id int64) error {
	return s.BillRepository.Delete(id)
}

func (s *BillDomainService) Update(bill entity.Bill) error {
	return s.BillRepository.Update(bill)
}

func (s *BillDomainService) GetById(id int64) *do.BillDo {
	return s.BillQuery.GetById(id)
}

func (s *BillDomainService) GetByUser(userName string, limit int, offset int) (total int64, list []do.BillDo) {
	return s.BillQuery.GetByUser(userName, limit, offset)
}
