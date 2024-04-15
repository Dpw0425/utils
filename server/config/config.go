package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"utils/server/logger"
)

var (
	CONFIG  Config
	MYSQLDB *gorm.DB
)

func InitConfig() {
	viper.SetConfigFile("./server/config/config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("读取配置失败: %v", err)
		return
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		logger.Error("解析配置失败: %v", err)
		return
	}

	logger.Info("读取配置成功！")
}
