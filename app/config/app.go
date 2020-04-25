package config

import (
	"math/rand"
	"strconv"
	"time"
)

/**
 * app配置信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func app() map[string]interface{} {
	/* 获取环境、版本信息 */
	version := env("app.version", nil)
	envL := env("app.env", "")
	if envL == "development" && version != nil {
		rand.Seed(time.Now().Unix())
		version = version.(string) + "__" + strconv.Itoa(rand.Intn(1000000))
	}
	return map[string]interface{}{
		"name":      "app",
		"http_port": env("app.http_port", nil),
		"version":   version,
		"env":       envL,
		"mysql": map[string]interface{}{
			"host":     env("mysql.host", ""),
			"port":     env("mysql.port", nil),
			"database": env("mysql.database", nil),
			"user":     env("mysql.user", nil),
			"password": env("mysql.password", nil),
		},
	}
}
