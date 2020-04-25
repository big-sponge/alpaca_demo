package models

type Config struct {
	Base
	Key   string `gorm:"column:key"`
	Value string `gorm:"column:value"`
}

/**
 * 设置表名
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func (Config) TableName() string {
	return "tb_config"
}
