package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	baseDir, _ := filepath.Abs(filepath.Dir(""))
	logFilePath := baseDir + "/log/"
	logFileName := "logs.txt"

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Create(fileName)
		src, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	//设置日志时间格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}
}

