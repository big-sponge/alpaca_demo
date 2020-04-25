package common

import (
	"alpaca_demo/app/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
 * 初始化Mysql
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Mysql() *gorm.DB {
	host := config.GetString("app.mysql.host")
	port := config.GetString("app.mysql.port")
	user := config.GetString("app.mysql.user")
	password := config.GetString("app.mysql.password")
	database := config.GetString("app.mysql.database")
	conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println(conn, err)
		panic(err)
	}
	db.LogMode(true)
	return db
}
