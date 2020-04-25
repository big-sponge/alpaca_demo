package v1

import (
	"alpaca_demo/app/api"
	"alpaca_demo/app/common/code"
	"alpaca_demo/app/models"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
 * 编辑学生信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func EditStudent(c *gin.Context) {
	/* 获取输入参数 */
	input := map[string]interface{}{
		"Id":      api.Input(c, "Id", ""),
		"Name":    api.Input(c, "Name", ""),
		"ClassId": api.Input(c, "ClassId", nil),
	}

	/* 校验参数 */
	if input["Name"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".Name"))
		return
	}
	if input["Number"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".Number"))
		return
	}
	if input["ClassId"] == nil {
		api.Output(c, nil, errors.New(code.FailParamNull+".ClassId"))
		return
	}

	/* 调用服务 */
	res := models.EditStudent(input)

	/* 返回结果 */
	api.Output(c, res, nil)
	return
}

/**
 * 学生信息列表
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ListStudent(c *gin.Context) {
	/* 获取输入参数 */
	input := map[string]interface{}{
		"Id":      api.Input(c, "Id", ""),
		"Name":    api.Input(c, "Name", ""),
		"ClassId": api.Input(c, "ClassId", nil),
	}

	/* 调用服务 */
	list, total := models.ListStudent(input)
	res := map[string]interface{}{
		"list":  list,
		"total": total,
	}

	/* 返回结果 */
	api.Output(c, res, nil)
	return
}

/**
 * 编辑学生信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func EditScore(c *gin.Context) {
	/* 获取输入参数 */
	input := map[string]interface{}{
		"Id":      api.Input(c, "Id", ""),
		"Name":    api.Input(c, "Name", ""),
		"ClassId": api.Input(c, "ClassId", nil),
	}

	/* 校验参数 */
	if input["Name"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".Name"))
		return
	}
	if input["Number"] == "" {
		api.Output(c, nil, errors.New(code.FailParamNull+".Number"))
		return
	}
	if input["ClassId"] == nil {
		api.Output(c, nil, errors.New(code.FailParamNull+".ClassId"))
		return
	}

	/* 调用服务 */
	res := models.EditStudent(input)

	/* 返回结果 */
	api.Output(c, res, nil)
	return
}

/**
 * 成绩信息列表
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ListScore(c *gin.Context) {
	/* 获取输入参数 */
	input := map[string]interface{}{
		"Id":        api.Input(c, "Id", nil),
		"Course":    api.Input(c, "Course", nil),
		"StudentId": api.Input(c, "StudentId", nil),
	}

	/* 调用服务 */
	list, total := models.ListScore(input)
	res := map[string]interface{}{
		"list":  list,
		"total": total,
	}

	/* 返回结果 */
	api.Output(c, res, nil)
	return
}
