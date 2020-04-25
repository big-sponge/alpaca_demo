package admin

import (
	"alpaca_demo/app/api"
	"alpaca_demo/app/common/code"
	"github.com/gin-gonic/gin"
)

/**
 * 查看当前登录用户的信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Info(c *gin.Context) {
	/* 获取信息 */
	member := api.InputVisitor(c, "LoginUser", nil)
	result := map[string]interface{}{
		"member": member,
	}
	/* 返回结果 */
	api.OutputFull(c, result, nil)
	return
}

/**
 * 修改密码
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ResetPwdByOld(c *gin.Context) {
	result := map[string]interface{}{
		"code": code.Success,
		"data": map[string]interface{}{
			"member": map[string]interface{}{
				"username": "admin",
				"name":     "admin",
			},
		},
	}
	/* 返回结果 */
	api.OutputFull(c, result, nil)
	return
}
