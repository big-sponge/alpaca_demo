package routers

import (
	"alpaca_demo/app/api"
	"alpaca_demo/app/common"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
 * 中间件处理Recovery
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Abort()
				api.OutputFull(c, nil, errors.New(common.Recovery()))
			}
		}()
		c.Next()
	}
}
