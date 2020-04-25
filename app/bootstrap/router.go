package bootstrap

import (
	"alpaca_demo/app/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/**··
 * 初始化路由
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitRouter() {

	// 创建路由
	r := gin.Default()

	// 静态页面
	r.StaticFS("/web", gin.Dir("./web", true))

	// 初始化路由
	routers.InitRouter(r)

	// 初始化管理接口路由
	routers.InitAdminRouter(r)

	// 启动web监听
	port := ":80"
	if viper.Get("app.http_port") != nil {
		port = ":" + viper.GetString("app.http_port")
	}
	_ = r.Run(port)
}
