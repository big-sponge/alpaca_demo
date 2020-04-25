package models

import (
	"alpaca_demo/app/common/code"
	"errors"
)

type AdminMember struct {
	Base
	UserName string `gorm:"column:username"`
	PassWd   string `gorm:"column:passwd"`
	Name     string `gorm:"column:name"`
	Role     string `gorm:"column:role"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (AdminMember) TableName() string {
	return "tb_admin_member"
}

/**
 * Lists
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ListsAdminMember(input map[string]interface{}) (models []AdminMember, total int) {

	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()

	//查询条件
	if input["Id"] != nil {
		db = db.Where("id = ?", input["Id"])
	}
	if input["UserName"] != nil {
		db = db.Where("username = ?", input["UserName"])
	}

	//查询数量
	db.Model(&AdminMember{}).Count(&total)

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
func EditAdminMember(input map[string]interface{}) (model AdminMember) {

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
	if input["UserName"] != nil {
		model.UserName = input["UserName"].(string)
	}
	if input["Name"] != nil {
		model.Name = input["Name"].(string)
	}
	if input["PassWd"] != nil {
		model.PassWd = input["PassWd"].(string)
	}
	if input["Role"] != nil {
		model.Role = input["Role"].(string)
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
func DeleteAdminMember(input map[string]interface{}) {
	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()

	if input["Id"] != nil {
		db.Delete(AdminMember{}, "id = ?", input["Id"])
	}
}

/**
 * Detail
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func DetailAdminMember(input map[string]interface{}) (AdminMember, error) {
	// 查询实例信息
	list, _ := ListsAdminMember(input)
	if len(list) == 0 {
		return AdminMember{}, errors.New(code.FailDataNotFount)
	}
	// 返回结果
	return list[0], nil
}
