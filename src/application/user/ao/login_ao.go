package ao

import "github.com/marshhu/ma-frame/utils"

//登录实体
type LoginAo struct {
	UserName string `json:"nameName"`
	Password string `json:"password"`
}

func (input *LoginAo) Validate() (string, bool) {
	if utils.IsEmpty(input.UserName) {
		return "请输入用户名", false
	}
	if utils.IsEmpty(input.Password) {
		return "请输入密码", false
	}
	return "", true
}
