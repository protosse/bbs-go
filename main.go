package main

import "C"
import (
	"bbs-go/app"
	"bbs-go/common/config"
	"bbs-go/util/logging"
	"flag"
	"fmt"
)

var (
	configFile = flag.String("config", config.DefaultPath, "配置文件路径")
)

func init() {
	flag.Parse()
	var err error

	// 初始化配置信息
	err = config.Init(*configFile)
	if err != nil {
		fmt.Printf("init config err: %v", err)
		return
	}

	// 初始化日志
	logging.Init()

	// 初始化数据库
	err = app.NewGormServer(config.Global).Connect()
	if err != nil {
		logging.Errorf("connect db failed: %v", err)
	}
}

func main() {
	app.NewIrisServer(config.Global).Run()
}
