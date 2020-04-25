package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InitDaemon() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		s := <-c
		fmt.Println("stop ok", s)
		os.Exit(0)
	}()
	fmt.Println("start ok")
}
