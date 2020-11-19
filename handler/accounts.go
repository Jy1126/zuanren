package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	srv "zuanren/service"
)

type AddContentRequest struct {
	CN        string `json:"cn"`
	US        string `json:"us"`
	JP        string `json:"jp"`
	KR        string `json:"kr"`
	PH        string `json:"ph"`
	Other     string `json:"other"`
	OtherNote string `json:"otherNote"`
}

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

func GetContent(c *gin.Context) {
	originCode := c.DefaultQuery("originCode", "cn")
	targetCode := c.DefaultQuery("targetCode", "")
	content := srv.GetContent(originCode, targetCode)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": gin.H{
			"items": content,
		},
	})

}

func AddContent(c *gin.Context) {
	json := AddContentRequest{}

	_ = c.BindJSON(&json)
	_ = srv.AddContent(json.CN, json.US, json.JP, json.KR, json.PH, json.Other, json.OtherNote)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": gin.H{
		},
	})

}
