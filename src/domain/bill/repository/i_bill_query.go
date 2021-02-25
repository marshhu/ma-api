package repository

import "github.com/marshhu/ma-api/src/domain/bill/do"

type IBillQuery interface {
	GetById(id int64) *do.BillDo
	GetByUser(userName string, limit int, offset int) (total int64, list []do.BillDo)
}
