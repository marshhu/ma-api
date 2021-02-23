package ao

//应用层实体对象
type UserAo struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	Gender     int64  `json:"gender"`
	CreateTime string `json:"createTime" timeFormat:"2006-01-02 15:04:05"`
	UpdateTime string `json:"updateTime" timeFormat:"2006-01-02 15:04:05"`
}
