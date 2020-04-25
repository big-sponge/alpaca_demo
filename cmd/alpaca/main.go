package main

import (
	"alpaca_demo/app/bootstrap"
	"alpaca_demo/app/common"
	"fmt"
)

func init() {
	/* 初始化命令行 */
	bootstrap.InitFlag()
	/* 初始化配置文件 */
	bootstrap.InitConfig()
	/* 配置PID */
	bootstrap.InitPid()
	/* 配置Log */
	bootstrap.InitLog()
	/* 配置进程 */
	bootstrap.InitDaemon()
	/* 最后启动http路由，同时用来阻塞主进程 */
	bootstrap.InitRouter()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("主程序异常: ", err, ",Recovery:", common.Recovery())
		}
	}()
}
