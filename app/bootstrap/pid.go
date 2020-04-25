package bootstrap

import (
	"alpaca_demo/app/common"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/**
 * 写入pidFile
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitPid() {
	pidFile := common.BasePath + "/storage/pidfile/pidfile"
	pidStr := strconv.Itoa(os.Getppid())
	pid := []byte(pidStr)
	if ioutil.WriteFile(pidFile, pid, 0644) == nil {
		fmt.Println("write pid successful:", pidFile, pidStr)
	}
}
