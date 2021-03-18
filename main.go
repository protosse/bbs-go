package main

import "C"
import (
	"bbs-go/app"
	"bbs-go/common/config"
	"bbs-go/models"
	"bbs-go/services"
	"bbs-go/util/logging"
	"flag"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	configFile = flag.String("config", "./config.yaml", "配置文件路径")
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

	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Config.DB.Prefix,
			SingularTable: true,
		},
	}

	// 连接数据库
	if err = services.OpenDB(config.Config.DB.Conn, gormConf, 10, 20, models.Models...); err != nil {
		logging.Errorf("connect db failed: %v", err)
	}
}

func main() {
	app.InitIris()
}
