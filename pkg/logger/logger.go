package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObject *logrus.Logger

func InitLog() {
	if LogrusObject != nil {
		src, _ := setOutputFile()
		//设置输出
		LogrusObject.Out = src
	}
	logger := logrus.New() //实例化
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := "" //设置日志文件路径
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	//检查日志文件路径是否存在,如果不存在则创建
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//构造日志文件名
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件完整路径
	fileName := path.Join(logFilePath, logFileName)
	//检查日志文件是否存在,如果不存在则创建
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 打开日志文件，准备写入日志
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
