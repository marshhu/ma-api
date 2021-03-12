package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/infrastructure/jwt"
	"github.com/marshhu/ma-api/src/application/user/ao"
	"github.com/marshhu/ma-api/src/interface/api"
	"github.com/marshhu/ma-api/src/interface/ioc"
	"github.com/marshhu/ma-frame/log"
	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"
)

// @Summary 注册
// @Description register
// @Tags  user
// @Accept  json
// @Produce  json
// @Param input body ao.RegisterAo true " "
// @Success 200 "ok"
// @Router /user/register [post]
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := ao.RegisterAo{}
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			api.BadRequestError("请求参数解析异常", ctx)
			return
		}
		if msg, ok := input.Validate(); !ok {
			api.BadRequestError(msg, ctx)
			return
		}

		id, err := ioc.DIContainer.UserAppService.Add(input)
		if err != nil {
			api.Ok(api.Failed(-1, "注册用户失败"), ctx)
			log.Errorf("注册用户发生异常:%s", err)
			return
		}
		api.Ok(api.Success(id), ctx)
	}
}

// @Summary 用户信息
// @Description user info
// @Tags  user
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Success 200 "ok"
// @Router /user/info [get]
func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		var userId int64
		if utils.IsEmpty(token) {
			user := api.GetUserInfo(ctx)
			if user.UserId <= 0 {
				api.Unauthorized("获取用户信息失败，请登录", ctx)
				return
			}
			userId = user.UserId
		} else {
			mc, err := jwt.ParseToken(token, viper.GetString("jwt.secret"))
			if err != nil {
				api.BadRequestError("token格式错误", ctx)
				return
			}
			userId = mc.UserId
		}
		info := ioc.DIContainer.UserAppService.GetById(userId)
		result := ao.UserInfoAo{ID: info.ID, Name: info.Name, NickName: info.NickName, Avatar: info.Avatar, Gender: info.Gender}
		api.Ok(api.Success(result), ctx)
	}
}
