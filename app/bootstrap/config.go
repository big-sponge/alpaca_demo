package bootstrap

import (
	"alpaca_demo/app/common"
	"alpaca_demo/app/config"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

/**
 * 初始化配置文件
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitConfig() {

	/* 配置文件的路径,读取环境变量 */
	env := os.Getenv("ENV_MODE")
	if env == "" {
		env = "development"
	}
	baseDir := os.Getenv("BASE_DIR")
	if baseDir == "" {
		baseDir = "."
	}
	fmt.Println("BASE_DIR:" + baseDir)
	fmt.Println("ENV:" + env)

	/* 配置文件的名字 */
	viper.SetConfigName(env)

	/* 配置文件的类型 */
	viper.SetConfigType("json")

	/* 设置根目录 */
	common.BasePath = baseDir

	/* 读取配置文件 */
	viper.AddConfigPath(common.BasePath + "/env")

	/* 读取配置文件信息 */
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	/* 设置代码配置文件 */
	config.InitConfig()
}
