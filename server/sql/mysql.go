package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"utils/server/config"
	"utils/server/logger"
)

func InitMysql() {
	m := config.CONFIG.Mysql
	db, err := gorm.Open(mysql.Open(config.DSN(m)), &gorm.Config{})
	if err != nil {
		logger.Error("连接 Mysql 失败: %v", err)
	} else {
		// TODO: auto migrate

		config.MYSQLDB = db
		logger.Info("连接 Mysql 成功: %v", config.DSN(m))
	}
}
