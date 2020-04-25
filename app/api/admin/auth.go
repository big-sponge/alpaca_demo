package admin

import (
	"alpaca_demo/app/api"
	"alpaca_demo/app/common/code"
	"alpaca_demo/app/service/admin"
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

/**
 * 登录
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Login(c *gin.Context) {
	/* 获取输入参数 */
	input := map[string]interface{}{
		"UserName": api.Input(c, "UserName", ""),
		"PassWd":   api.Input(c, "PassWd", ""),
	}

	/*检查参数*/
	if input["UserName"] != nil && reflect.TypeOf(input["UserName"]).String() != "string" {
		api.Output(c, nil, errors.New(code.FailParamFormat+".UserName"))
		return
	}
	if input["PassWd"] != nil && reflect.TypeOf(input["PassWd"]).String() != "string" {
		api.Output(c, nil, errors.New(code.FailParamFormat+".PassWd"))
		return
	}
	if input["UserName"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".UserName"))
		return
	}
	if input["PassWd"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".PassWd"))
		return
	}

	/* 获取信息 */
	res, err := admin.CheckLogin(input)
	if err != nil {
		api.Output(c, res, err)
		return
	}

	/*设置cookie*/
	cookieToken := res.(map[string]interface{})["token"].(string)
	c.SetCookie("admin_token", cookieToken, 3600, "/", "", false, true)

	/* 返回结果 */
	api.Output(c, res, err)
	return
}

/**
 * 注销
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Logout(c *gin.Context) {
	/* 获取输入参数 */
	c.SetCookie("admin_token", "", -1, "/", "", false, true)
	/* 返回结果 */
	api.Output(c, nil, nil)
	return
}
