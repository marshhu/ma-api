package repository

import "github.com/marshhu/ma-api/src/domain/bill/entity"

type IBillRepository interface {
	Add(bill entity.Bill) (id int64, err error)
	Delete(id int64) error
	Update(bill entity.Bill) error
}
