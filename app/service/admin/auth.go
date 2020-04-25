package admin

import (
	"alpaca_demo/app/common"
	"alpaca_demo/app/common/code"
	"alpaca_demo/app/models"
	"fmt"
	"github.com/pkg/errors"
)

/**
 * 实例状态列表
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func GetLoginUserInfo(token string) (result interface{}, err error) {
	sessionInfo := models.GetAdminSession(token)
	if sessionInfo == nil {
		return nil, err
	}
	userInfo, err := models.DetailAdminMember(map[string]interface{}{
		"Id": sessionInfo.(models.AdminSession).MemberId,
	})
	if err != nil {
		return nil, err
	}
	result = map[string]interface{}{
		"Id":       userInfo.Id,
		"UserName": userInfo.UserName,
		"Name":     userInfo.Name,
		"Role":     userInfo.Role,
	}
	return result, err
}

/**
 * 检查用户名密码
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func CheckLogin(input map[string]interface{}) (result interface{}, err error) {

	/* 查找用户信息 验证用户名*/
	userInfo, err := models.DetailAdminMember(map[string]interface{}{
		"UserName": input["UserName"],
	})
	if err != nil {
		return result, errors.New(code.FailParamError + ".UserName")
	}

	/* 查找用户密码 */
	if !common.PasswordVerify(userInfo.PassWd, input["PassWd"].(string)) {
		return result, errors.New(code.FailParamError + ".PassWd")
	}

	/* 生成token */
	token := common.InitToken(fmt.Sprintf("%d", userInfo.Id))

	/* 保存session信息 */
	models.EditAdminSession(map[string]interface{}{
		"Token":    token,
		"MemberId": userInfo.Id,
	})

	/* 查找用户信息 */
	result = map[string]interface{}{
		"member": map[string]interface{}{
			"username": userInfo.UserName,
			"name":     userInfo.Name,
		},
		"token": token,
	}
	return result, err
}
