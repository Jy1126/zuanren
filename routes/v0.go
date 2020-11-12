package routes

import (
	"github.com/gin-gonic/gin"
	"xtp_account_check/handler"
)

func SetupV0Router(r *gin.Engine) {

	// v0
	v0Group := r.Group("v0")
	{
		// 绑定账号数
		// 查询
		v0Group.GET("/bindCount", handler.GetBindCount)
	}
}
