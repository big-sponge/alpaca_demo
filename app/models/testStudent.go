package models

import "time"

type Student struct {
	Base
	Name     string     `gorm:"column:name"`
	Number   string     `gorm:"column:number"`
	Gender   int        `gorm:"column:gender"`
	Birthday *time.Time `gorm:"column:birthday"`
	ClassId  uint       `gorm:"column:class_id"`
	Class    Class      `gorm:"ForeignKey:ClassId;AssociationForeignKey:Id"`
	Scores   []Score    `gorm:"ForeignKey:Id;AssociationForeignKey:StudentId"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (Student) TableName() string {
	return "tb_test_student"
}

/**
 * 保存
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (m *Student) Save() Student {
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()
	db = db.Save(m)
	if db.Error != nil {
		panic(db.Error)
	}
	return *m
}

/**
 * 删除
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (m *Student) Delete() {
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()
	db = db.Delete(m)
	if db.Error != nil {
		panic(db.Error)
	}
}

/**
 * Lists
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ListStudent(input map[string]interface{}) (models []Student, total int) {

	/* 连接数据库 */
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()

	//查询条件
	if input["Id"] != nil {
		db = db.Where("id = ?", input["Id"])
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
func EditStudent(input map[string]interface{}) (model Student) {

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
	if input["Name"] != nil {
		model.Name = input["Name"].(string)
	}
	if input["Number"] != nil {
		model.Number = input["Number"].(string)
	}
	if input["Birthday"] != nil {
		model.Birthday = input["Birthday"].(*time.Time)
	}
	if input["Birthday"] != nil {
		model.ClassId = input["ClassId"].(uint)
	}
	if input["Gender"] != nil {
		model.Gender = input["Gender"].(int)
	}

	// 保存
	r := db.Save(&model)
	if r.Error != nil {
		panic(r.Error)
	}
	return model
}
