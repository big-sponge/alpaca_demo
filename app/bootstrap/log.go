package bootstrap

import (
	"alpaca_demo/app/common"
	"alpaca_demo/app/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

/**
 * 初始化配置log
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitLog() {
	// log文件位置
	baseLogPath := config.GetString("log.dir") + config.GetString("log.filename")
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		// 文件最大保存时间
		rotatelogs.WithMaxAge(config.Get("log.max_age").(time.Duration)),
		// 日志切割时间间隔
		rotatelogs.WithRotationTime(config.Get("log.rotation_time").(time.Duration)),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 设置gin框架的gin
	gin.DefaultWriter = writer

	// 设置一般Log
	common.Log = logrus.New()
	common.Log.SetOutput(writer)
	common.Log.SetFormatter(&logrus.JSONFormatter{})
	common.Log.SetLevel(logrus.ErrorLevel | logrus.InfoLevel | logrus.DebugLevel)

	// 设置系统错误Log
	common.LogError = logrus.New()
	common.LogError.SetOutput(writer)
	common.LogError.SetFormatter(&logrus.JSONFormatter{})
	common.LogError.SetLevel(logrus.ErrorLevel | logrus.InfoLevel | logrus.DebugLevel)
}
