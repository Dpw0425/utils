package server

import (
	"github.com/gin-gonic/gin"
	"utils/server/config"
	"utils/server/logger"
	"utils/server/sql"
)

func InitServer() *gin.Engine {
	logger.InitZap()
	config.InitConfig()
	sql.InitMysql()

	r := gin.Default()

	return r
}
