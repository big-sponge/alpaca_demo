package api

import (
	"alpaca_demo/app/common"
	"alpaca_demo/app/common/code"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

/**
 * 获取会话数据
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InputVisitor(c *gin.Context, key string, defaultValue interface{}) interface{} {
	value, exists := c.Get(key)
	if exists == true {
		return value
	}
	return defaultValue
}

/**
 * 封装获取输入参数
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Input(c *gin.Context, key string, defaultValue interface{}) interface{} {

	// 初始化
	var isGet bool
	value := defaultValue

	// 先获取path参数
	value = c.Param(key)
	if value != "" {
		return value
	}

	// 先获取GET参数
	value, isGet = c.GetQuery(key)
	if isGet == true {
		return value
	}

	// 获取POST参数
	value, isGet = c.GetPostForm(key)
	if isGet == true {
		return value
	}

	// 获取json数据
	var json interface{}
	err := c.ShouldBindBodyWith(&json, binding.JSON)
	if err == nil && json != nil {
		if json.(map[string]interface{})[key] != nil {
			value = json.(map[string]interface{})[key]
			return value
		}
	}

	// 返回结果
	return defaultValue
}

/**
 * 封装获取输入参数Map
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InputMap(c *gin.Context, key string, defaultValue interface{}) interface{} {

	// 初始化
	var err bool
	value := defaultValue

	// 先获取GET参数
	value, err = c.GetQueryMap(key)
	if err == false {
		// 若GET不存在，获取POST参数
		value, err = c.GetPostFormMap(key)
		if err == false {
			// 都不存在使用默认值
			value = defaultValue
		}
	}

	// 返回结果
	return value
}

/**
 * 封装获取输入参数Cookie
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InputCookie(c *gin.Context, key string, defaultValue interface{}) interface{} {

	// 初始化
	value := defaultValue

	// 先获取Cookie参数
	value, err := c.Cookie(key)
	if err != nil {
		value = defaultValue
	}

	// 返回结果
	return value
}

/**
 * 封装获取输入参数Header
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InputHeader(c *gin.Context, key string, defaultValue interface{}) interface{} {

	// 初始化
	value := defaultValue

	// 先获取GET参数
	value = c.GetHeader(key)

	// 返回结果
	return value
}

/**
 * 返回输出格式化
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func OutputFull(c *gin.Context, data interface{}, err error) {

	/* 返回正确结果 */
	if err == nil {
		res := gin.H{"code": code.Success}
		if data != nil {
			res["data"] = data
		}
		c.JSON(200, res)
		return
	}

	/* 用户错误 */
	if strings.HasPrefix(err.Error(), code.Fail) == true {
		res := gin.H{"code": err.Error()}
		if data != nil {
			res["data"] = data
		}
		c.JSON(200, res)
		return
	}

	/* 系统错误 */
	c.JSON(200, gin.H{"code": code.Error, "error": err.Error()})
	common.Log.Error(err.Error())
	return
}
