package bootstrap

import (
	"flag"
)

var (
	ShowVersion bool
	Test        string
)

/**
 * 初始化flag
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func InitFlag() {
	flag.BoolVar(&ShowVersion, "version", false, "Print version information.")
	flag.StringVar(&Test, "test", "test string", "test.")
	flag.Parse()
}
