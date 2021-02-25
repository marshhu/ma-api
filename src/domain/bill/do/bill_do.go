package do

//领域层实体对象
type BillDo struct {
	ID          int64   `json:"id"`
	Type        int     `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	BillDate    string  `json:"billDate" timeFormat:"2006-01-02"`
	CreateTime  string  `json:"createTime" timeFormat:"2006-01-02 15:04:05"`
	UpdateTime  string  `json:"updateTime" timeFormat:"2006-01-02 15:04:05"`
	CreateBy    string  `json:"createBy"`
}
