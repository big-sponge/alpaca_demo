package routers

import (
	baseApi "alpaca_demo/app/api"
	adminApi "alpaca_demo/app/api/admin"
	"alpaca_demo/app/common/code"
	"alpaca_demo/app/service/admin"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
 * 中间件，判断用户是否登录
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		/* 获取输入参数 */
		cookieToken := baseApi.InputCookie(c, "admin_token", nil)
		token := baseApi.Input(c, "admin_token", nil)

		/* 检查token是否存在 */
		if token == nil {
			token = cookieToken
		}
		if token == nil {
			c.Abort()
			baseApi.OutputFull(c, nil, errors.New(code.FailLogin))
			return
		}

		/* 获取获取用户信息 */
		userInfo, err := admin.GetLoginUserInfo(token.(string))
		if err != nil || userInfo == nil {
			/* 返回结果 */
			c.Abort()
			baseApi.OutputFull(c, nil, errors.New(code.FailLogin))
			return
		}

		/* 保存用户信息 */
		c.Set("LoginUser", userInfo)

		/* 验证通过 */
		c.Next()
		return
	}
}

/**
 * 初始化Admin模块路由
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitAdminRouter(r *gin.Engine) {

	auth := r.Group("/admin/auth/")
	auth.Use(Recovery())
	auth.Any("/login", adminApi.Login)   /* 登录接口 */
	auth.Any("/logout", adminApi.Logout) /* 注销接口 */

	router := r.Group("/admin")
	router.Use(Recovery())                                    /* Recovery */
	router.Use(AdminAuthMiddleware())                         /* 登录验证中间件，判断用户是否登录 */
	router.Any("/self/info", adminApi.Info)                   /* Info */
	router.Any("/self/resetPwdByOld", adminApi.ResetPwdByOld) /* 修改密码 */
}
