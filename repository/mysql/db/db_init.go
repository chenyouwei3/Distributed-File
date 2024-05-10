package db

import (
	"Search-Engine/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB

func InitMySQL() {
	mConfig := config.Conf.MySQL
	dsn := strings.Join([]string{
		mConfig.UserName, ":",
		mConfig.Password, "@tcp(",
		mConfig.Host, ":",
		mConfig.Port, ")/",
		mConfig.Database,
		"?charset=" + mConfig.Charset + "&parseTime=true"},
		"")
	database(dsn)
}

func database(connString string) error {
	var ormLogger logger.Interface
	// 根据 Gin 的运行模式设置日志记录器
	if gin.Mode() == "debug" {
		//如果是 debug 模式，则使用 Info 级别记录日志，
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		//如果是 debug 模式，则使用 Info 级别记录日志，
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // 数据源名称
		DefaultStringSize:         256,        // 字符串类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 设置 GORM 的日志记录器
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名策略
		},
	})
	if err != nil {
		err = errors.Wrap(err, "failed to open Mysql")
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migration()
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	return _db.WithContext(ctx)
}
