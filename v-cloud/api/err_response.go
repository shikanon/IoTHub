package api

import "github.com/gin-gonic/gin"

func ErrResponse(msg string, c *gin.Context) {
	resp := gin.H{
		"status":  "N",
		"message": msg,
		"data":    nil,
	}
	c.JSON(400, resp)
}

func DbErrorResponse(msg string, c *gin.Context) {
	resp := gin.H{
		"status":  "N",
		"message": "操作失败",
		"data":    nil,
	}
	c.JSON(400, resp)
}

func AnalyticParameterErrResponse(c *gin.Context) {
	resp := gin.H{
		"status":  "N",
		"message": "参数解析错误",
		"data":    nil,
	}
	c.JSON(400, resp)
}
