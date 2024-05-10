package db

import (
	log "Search-Engine/pkg/logger"
	"Search-Engine/repository/mysql/model"
	"os"
)

func migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.LogrusObject.Infoln("register table fail")
		os.Exit(0)
	}
	log.LogrusObject.Infoln("register table success")
}
