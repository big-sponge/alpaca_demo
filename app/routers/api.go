package routers

import (
	"alpaca_demo/app/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * 初始化路由
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitRouter(r *gin.Engine) {

	/* 路由 */
	router := r.Group("")
	router.Use(Recovery())
	router.Any("/ping", v1.Ping)
	router.Any("/student/list", v1.ListStudent)
	router.Any("/student/edit", v1.EditStudent)
	router.Any("/score/list", v1.ListScore)
	router.Any("/score/edit", v1.EditScore)

}
