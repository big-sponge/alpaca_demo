package models

type Class struct {
	Base
	Name string `gorm:"column:name"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (Class) TableName() string {
	return "tb_test_class"
}

/**
 * 保存
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (m *Class) Save() Class {
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
func (m *Class) Delete() {
	db := Mysql()
	defer func() {
		_ = db.Close()
	}()
	db = db.Delete(m)
	if db.Error != nil {
		panic(db.Error)
	}
}
