package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/infrastructure/jwt"
	"github.com/marshhu/ma-api/src/application/user/ao"
	"github.com/marshhu/ma-api/src/interface/api"
	"github.com/marshhu/ma-api/src/interface/ioc"
	"github.com/spf13/viper"
)

// @Summary 登录
// @Description login
// @Tags  auth
// @Accept  json
// @Produce  json
// @Param input body ao.LoginAo true " "
// @Success 200 "ok"
// @Router /auth/login [post]
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := ao.LoginAo{}
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			api.BadRequestError("请求参数解析异常", ctx)
			return
		}
		if msg, ok := input.Validate(); !ok {
			api.BadRequestError(msg, ctx)
			return
		}

		user := ioc.DIContainer.UserAppService.Get(input.UserName, input.Password)
		if user.ID > 0 {
			token, err := jwt.GenToken(user.ID, user.Name, viper.GetString("jwt.secret"),
				viper.GetInt("jwt.expireTime"))
			if err != nil {
				api.Ok(api.Failed("生成token发生异常"), ctx)
				return
			}
			var userInfo = ao.UserInfoAo{
				ID:       user.ID,
				Name:     user.Name,
				NickName: user.NickName,
				Gender:   user.Gender,
				Avatar:   user.Avatar,
				Token:    token,
			}
			api.Ok(api.Success(userInfo), ctx)
			return
		}
		api.Ok(api.Failed("用户名或密码错误"), ctx)
	}
}
