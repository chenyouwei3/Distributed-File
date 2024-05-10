package loading

import (
	"Search-Engine/config"
	log "Search-Engine/pkg/logger"
	"Search-Engine/repository/mysql/db"
)

func Loading() {
	config.InitConfig() //配置文件
	log.InitLog()       //日志
	db.InitMySQL()      //数据库
}
