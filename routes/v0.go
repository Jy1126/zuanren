package routes

import (
	"github.com/gin-gonic/gin"
	"zuanren/handler"
)

func SetupV0Router(r *gin.Engine) {

	// v0
	v0Group := r.Group("v0")
	{
		// 绑定账号数
		// 查询
		v0Group.GET("/bindCount", handler.GetBindCount)

		// 内容
		// 查询
		v0Group.GET("/content", handler.GetContent)
		// 新增
		v0Group.POST("/content", handler.AddContent)
	}
}
