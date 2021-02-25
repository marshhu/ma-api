package repository

import (
	"github.com/marshhu/ma-api/src/domain/bill/entity"
	"github.com/marshhu/ma-frame/orm"
)

type BillRepository struct {
}

func (r *BillRepository) Add(bill entity.Bill) (id int64, err error) {
	db := orm.NewQuery()
	err = db.Create(&bill)
	if err != nil {
		return 0, err
	}
	return bill.Id, nil
}

func (r *BillRepository) Delete(id int64) error {
	db := orm.NewQuery()
	return db.Delete(&entity.Bill{}, id)
}

func (r *BillRepository) Update(bill entity.Bill) error {
	db := orm.NewQuery()
	return db.Model(&entity.Bill{}).Where("id = ?", bill.Id).
		Select("amount", "type", "description", "bill_date", "update_time").Updates(bill)
}
