package main

import (
	"bbs-go/app"
	"bbs-go/config"
	"bbs-go/models"
	"bbs-go/services"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	configFile = flag.String("config", "./config.yaml", "配置文件路径")
	Conf       *config.Config
)

func init() {
	flag.Parse()
	Conf, err := config.Init(*configFile)
	if err != nil {
		fmt.Printf("init config err: %v", err)
		return
	}

	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Conf.DB.Prefix,
			SingularTable: true,
		},
	}

	// 初始化日志
	if file, err := os.OpenFile(Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		switch Conf.LogLevel {
		default:
			logrus.SetLevel(logrus.DebugLevel)
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Errorf("open logFile failed: %v", err)
	}

	// 连接数据库
	if err := services.OpenDB(Conf.DB.Conn, gormConf, 10, 20, models.Models...); err != nil {
		logrus.Errorf("connect db failed: %v", err)
	}
}

func main() {
	app.InitIris(Conf)
}
