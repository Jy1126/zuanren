package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	srv "xtpwidget/service"
)

func GetBindCount(c *gin.Context) {
	bindCount, unBindCount := srv.GetBindCount()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": gin.H{
			"bindCount":   bindCount,
			"unbindCount": unBindCount,
		},
	})
}
