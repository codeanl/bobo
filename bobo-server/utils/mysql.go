package utils

import (
	"bobo-server/config"
	"bobo-server/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func InitMySQL() *gorm.DB {
	mysqlCfg := config.Cfg.Mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm 日志模式
		Logger: logger.Default.LogMode(getLogMode(config.Cfg.Mysql.LogMode)),
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("MySQL 连接失败, 请检查参数")
	}
	log.Println("MySQL 连接成功")
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	MakeMigrate(db)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)                 // 设置连接池中的最大闲置连接
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置连接的最大可复用时间
	return db
}

// 根据字符串获取对应 LogLevel
func getLogMode(str string) logger.LogLevel {
	switch str {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

func MakeMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Daily{},
		&model.Topic{},
		&model.DailyTopic{},
		&model.DailyAttachment{},
		&model.Article{},
		&model.OperationLog{},
		&model.LoginLog{},
	)
	if err != nil {
		log.Println("gorm 自动迁移失败: ", err)
	} else {
		log.Println("gorm 自动迁移成功")
	}
}
