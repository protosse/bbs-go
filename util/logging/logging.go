package logging

import (
	"bbs-go/common/config"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	if file, err := os.OpenFile(config.Global.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		switch config.Global.LogLevel {
		default:
			logrus.SetLevel(logrus.DebugLevel)
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Errorf("open logFile failed: %v", err)
	}
}

func Debugf(msg string, args ...interface{}) {
	logrus.Debugf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	logrus.Errorf(msg, args...)
}
