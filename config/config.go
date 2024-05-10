package config

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	MySQL *MySQL `yaml:"mysql"`
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"userName"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config") //设置配置文件的文件名为”config“
	viper.SetConfigType("yaml")   //设置配置文件的格式名为”yaml“
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig() //读取配置文件
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
