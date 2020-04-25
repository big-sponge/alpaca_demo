package models

import (
	"alpaca_demo/app/common/code"
	"errors"
	"time"
)

type AdminSession struct {
	Base
	MemberId      uint       `gorm:"column:member_id"`
	Token         string     `gorm:"column:token"`
	LoginTime     *time.Time `gorm:"column:login_time"`
	UpdateTime    *time.Time `gorm:"column:update_time"`
	AvailableTime *time.Time `gorm:"column:available_time"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (AdminSession) TableName() string {
	return "tb_admin_session"
}

/**
 * Lists
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ListsAdminSession(input map[string]interface{}) (models []AdminSession, total int) {

	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()

	//查询条件
	if input["Id"] != nil {
		db = db.Where("id = ?", input["Id"])
	}
	if input["Token"] != nil {
		db = db.Where("token = ?", input["Token"])
	}

	//查询数量
	db.Model(&AdminSession{}).Count(&total)

	//排序
	db = InitOrdered(db, input["orders"])

	//分页
	db = InitPaged(db, input["page_size"], input["page_num"])

	//查找数据
	db.Find(&models)

	//返回结果
	return models, total
}

/**
 * Edit
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func EditAdminSession(input map[string]interface{}) (model AdminSession) {

	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()
	// 判断是否指定了Id。如果指定了，就是修改，否则是新增
	if input["Id"] != nil {
		r := db.First(&model, input["Id"])
		if r.Error != nil {
			panic(r.Error)
		}
	} else {
		r := db.Create(&model)
		if r.Error != nil {
			panic(r.Error)
		}
	}

	// 设置字段
	if input["Token"] != nil {
		model.Token = input["Token"].(string)
	}
	if input["MemberId"] != nil {
		model.MemberId = input["MemberId"].(uint)
	}
	if input["AvailableTime"] != nil {
		model.AvailableTime = input["AvailableTime"].(*time.Time)
	}

	// 保存
	r := db.Save(&model)
	if r.Error != nil {
		panic(r.Error)
	}
	return model
}

/**
 * Delete
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func DeleteAdminSession(input map[string]interface{}) {

	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()
	if input["Id"] != nil {
		db.Delete(AdminSession{}, "id = ?", input["Id"])
	}
	if input["Token"] != nil {
		db.Delete(AdminSession{}, "token = ?", input["Token"])
	}
	// 返回结果
	return
}

/**
 * Edit
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func DetailAdminSession(input map[string]interface{}) (AdminSession, error) {
	// 查询实例信息
	list, _ := ListsAdminSession(input)
	if len(list) == 0 {
		return AdminSession{}, errors.New(code.FailDataNotFount)
	}

	// 返回结果
	return list[0], nil
}

/**
 * 通过Token获取Session信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func GetAdminSession(token string) (result interface{}) {
	// 查询实例信息
	list, _ := ListsAdminSession(map[string]interface{}{
		"token": token,
	})
	if len(list) == 0 {
		return nil
	}
	// 返回结果
	return list[0]
}
