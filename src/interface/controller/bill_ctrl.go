package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/src/application/bill/ao"
	"github.com/marshhu/ma-api/src/interface/api"
	"github.com/marshhu/ma-api/src/interface/ioc"
	"github.com/marshhu/ma-frame/log"
	"github.com/marshhu/ma-frame/utils"
)

// @Summary 新增账单
// @Description add bill
// @Tags  bill
// @Accept  json
// @Produce  json
// @Param input body ao.AddBillAo true " "
// @Success 200 "ok"
// @Router /bill [post]
func AddBill() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := ao.AddBillAo{}
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			api.BadRequestError("请求参数解析异常", ctx)
			return
		}
		if msg, ok := input.Validate(); !ok {
			api.BadRequestError(msg, ctx)
			return
		}
		user := api.GetUserInfo(ctx)
		if user.UserId <= 0 {
			api.Unauthorized("获取用户信息失败，请登录", ctx)
			return
		}
		id, err := ioc.DIContainer.BillAppService.Add(input, user.UserName)
		if err != nil {
			log.Errorf("添加账单发生异常:%s", err)
			api.InternalServerError("添加账单失败", ctx)
		}
		api.Ok(id, ctx)
	}
}

// @Summary 获取账单列表
// @Description add bill
// @Tags  bill
// @Accept  json
// @Produce  json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 "ok"
// @Router /bill [get]
func GetBillByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit, _ := utils.ToInt(ctx.DefaultQuery("limit", "10"))
		offset, _ := utils.ToInt(ctx.DefaultQuery("offset", "0"))
		user := api.GetUserInfo(ctx)
		if user.UserId <= 0 {
			api.Unauthorized("获取用户信息失败，请登录", ctx)
			return
		}
		total, list := ioc.DIContainer.BillAppService.GetByUser(user.UserName, limit, offset)
		api.Ok(api.PageResult{Total: total, Data: list}, ctx)
	}
}
