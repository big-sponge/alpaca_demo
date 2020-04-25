package bootstrap

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

/**
 * 初始化Cron
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitCron() {
	i := 0
	c := cron.New()
	err := c.AddFunc("*/10 * * * * ?", func() {
		i = 1
		go func(){
			logrus.Info("test cron ", i)
		}()
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
}
