package models

import (
	"alpaca_demo/app/common"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

/**
 * Bass
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
type Base struct {
	Id        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

/**
 * Model
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (m *Base) Model() *Base {
	return m
}

/**
 * 连接数据库
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Mysql() *gorm.DB {
	return common.Mysql()
}

/**
 * 处理分页
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitPaged(db *gorm.DB, PageSize interface{}, PageNum interface{}) *gorm.DB {

	//设置limit
	if PageSize != nil {
		db = db.Limit(PageSize)
	}

	//设置offset
	if PageNum != nil {
		if PageSize == nil {
			PageSize = 0
		}
		num, _ := strconv.Atoi(PageNum.(string))
		size, _ := strconv.Atoi(PageSize.(string))
		offset := strconv.Itoa((num - 1) * size)
		db = db.Offset(offset)
	}
	return db
}

/**
 * 排序处理
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitOrdered(db *gorm.DB, Order interface{}) *gorm.DB {
	if Order == nil {
		return db
	}
	for k, v := range Order.(map[string]string) {
		if strings.ToUpper(v) != "ASC" {
			v = "DESC"
		}
		db = db.Order(k + " " + v)
	}
	return db
}
