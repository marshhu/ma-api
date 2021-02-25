package ao

import (
	"github.com/marshhu/ma-api/src/domain/bill/enums"
	"github.com/marshhu/ma-frame/utils"
)

type AddBillAo struct {
	Type        int     `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	BillDate    string  `json:"billDate"`
}

func (input *AddBillAo) Validate() (string, bool) {
	if input.Type != int(enums.Income) && input.Type != int(enums.Expenses) {
		return "账单类型错误", false
	}
	if input.Amount <= 0 {
		return "账单金额必须大于零", false
	}
	if !utils.MatchDate(input.BillDate) {
		return "账单日期格式错误", false
	}
	return "", true
}
