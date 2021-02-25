package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/src/application/bill/ao"
	"github.com/marshhu/ma-api/src/interface/api"
	"github.com/marshhu/ma-api/src/interface/ioc"
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

		id, err := ioc.DIContainer.BillAppService.Add(input)
		if err != nil {
			api.InternalServerError("添加账单失败", ctx)
		}
		api.Ok(id, ctx)
	}
}
