package models

type Score struct {
	Base
	Course    string    `gorm:"column:course"`
	StudentId uint      `gorm:"column:student_id"`
	Score     float64   `gorm:"column:score"`
	Students  []Student `gorm:"ForeignKey:StudentId;AssociationForeignKey:Id"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (Score) TableName() string {
	return "tb_test_score"
}

/**
 * 保存
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (m *Score) Save() Score {
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
func (m *Score) Delete() {
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
func ListScore(input map[string]interface{}) (models []Score, total int) {

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
	db.Model(&Score{}).Count(&total)

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
func EditScore(input map[string]interface{}) (model Score) {

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
	if input["StudentId"] != nil {
		model.StudentId = input["StudentId"].(uint)
	}
	if input["Course"] != nil {
		model.Course = input["Course"].(string)
	}
	if input["Score"] != nil {
		model.Score = input["Score"].(float64)
	}

	// 保存
	r := db.Save(&model)
	if r.Error != nil {
		panic(r.Error)
	}
	return model
}
