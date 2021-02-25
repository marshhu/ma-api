package repository

import (
	"github.com/marshhu/ma-api/src/domain/bill/do"
	"github.com/marshhu/ma-api/src/domain/bill/entity"
	"github.com/marshhu/ma-frame/orm"
	"github.com/marshhu/ma-frame/utils"
)

type BillQuery struct {
}

func (q *BillQuery) GetById(id int64) *do.BillDo {
	var bill entity.Bill
	db := orm.NewQuery()
	err := db.First(&bill, id).Error
	if err != nil {
		return nil
	}
	var result do.BillDo
	utils.MapTo(&bill, &result)
	return &result
}

func (q *BillQuery) GetByUser(userName string, limit int, offset int) (total int64, list []do.BillDo) {
	var bills []entity.Bill
	db := orm.NewQuery()
	err := db.Model(&entity.Bill{}).Where("create_by = ?", userName).
		Order("create_time desc").Offset(offset).Limit(limit).Find(&bills).Error
	if err != nil {
		return 0, nil
	}
	utils.MapTo(&bills, &list)
	db.Model(&entity.Bill{}).Count(&total)
	return total, list
}
