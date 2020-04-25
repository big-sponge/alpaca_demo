package config

import (
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

/**
 * 定义ValueConfig,存放配置信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
var ValueConfig map[string]interface{}

/**
 * 注册配置
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitConfig() {
	ValueConfig = map[string]interface{}{
		"app": app(),
		"log": log(),
	}
}

/**
 * 获取配置文件的内容
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func env(key string, dv interface{}) (val interface{}) {
	val = viper.Get(key)
	if val == nil {
		val = dv
	}
	return val
}

/**
 * 获取配置文件的内容 - Public
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Env(key string, dv interface{}) (val interface{}) {
	return env(key, dv)
}

/**
 * 根据Key获取内容
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Get(key string) (val interface{}) {
	keys := strings.Split(key, ".")
	kLen := len(keys)
	val = ValueConfig
	for i := 0; i < kLen; i++ {
		if val == nil {
			break
		}
		switch val.(type) {
		case map[string]interface{}:
			val = val.(map[string]interface{})[keys[i]]
		case []interface{}:
			index, err := strconv.Atoi(keys[i])
			if err != nil {
				return nil
			}
			val = val.([]interface{})[index]
		case interface{}:
			break
		default:
			break
		}
	}
	return val
}

/**
 * 根据Key获取内容 - string
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func GetString(key string) (val string) {
	return Get(key).(string)
}

/**
 * 根据Key获取内容 - int
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func GetInt(key string) (val int) {
	return Get(key).(int)
}

/**
 * 根据Key获取内容 - Duration
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func GetDuration(key string) (val time.Duration) {
	return Get(key).(time.Duration)
}
