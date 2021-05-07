package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/src/application/user/ao"
	"github.com/marshhu/ma-api/src/interface/api"
	"github.com/marshhu/ma-api/src/interface/ioc"
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
			api.Ok(api.Failed("注册用户失败"), ctx)
			return
		}
		api.Ok(api.Success(id), ctx)
	}
}
