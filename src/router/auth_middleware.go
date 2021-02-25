package router

import (
	"net/http"
	"strings"

	"github.com/marshhu/ma-api/infrastructure/jwt"
	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//debug环境下，不检查权限
		if viper.GetString("app.runMode") == "debug" {
			ctx.Next()
			return
		}
		if CheckAuth(ctx) {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 500,
			"msg":  "无权限访问，请登录",
		})
		ctx.Abort()
		return
	}
}

func CheckAuth(ctx *gin.Context) bool {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := ctx.Request.Header.Get("Authorization")
	if len(authHeader) == 0 {
		return false
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return false
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := jwt.ParseToken(parts[1], viper.GetString("jwt.secret"))
	if err != nil {
		return false
	}
	// 将当前请求的userName信息保存到请求的上下文ctx上
	ctx.Set("userId", utils.Int64ToString(mc.UserId))
	ctx.Set("userName", mc.UserName)
	return true
}
