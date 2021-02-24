package ao

import "github.com/marshhu/ma-frame/utils"

//注册实体
type RegisterAo struct {
	UserName        string `json:"nameName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (input *RegisterAo) Validate() (string, bool) {
	if utils.IsEmpty(input.UserName) {
		return "请输入用户名", false
	}
	if utils.IsEmpty(input.Password) {
		return "请输入密码", false
	}

	if utils.IsEmpty(input.ConfirmPassword) {
		return "请输入确认密码", false
	}

	if input.Password != input.ConfirmPassword {
		return "确认密码与原密码不一致", false
	}
	return "", true
}
