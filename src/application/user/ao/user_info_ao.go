package ao

type UserInfoAo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Gender   int64  `json:"gender"`
}
