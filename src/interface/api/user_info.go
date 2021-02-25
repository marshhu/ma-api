package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"
)

type UserInfo struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

func GetUserInfo(ctx *gin.Context) *UserInfo {
	userInfo := &UserInfo{}
	if viper.GetString("app.runMode") == "debug" {
		userInfo.UserId = viper.GetInt64("test.userId")
		userInfo.UserName = viper.GetString("test.userName")
		return userInfo
	}

	userId, err := utils.ToInt64(ctx.GetString("userId"))
	if err != nil {
		userInfo.UserId = 0
	} else {
		userInfo.UserId = userId
	}
	userInfo.UserName = ctx.GetString("userName")
	return userInfo
}
