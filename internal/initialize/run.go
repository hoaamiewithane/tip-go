package initialize

import (
	"fmt"
	"go-ecommerce/global"

	"go.uber.org/zap"
)

// Run all initialize
func Run() {
	// load config
	LoadConfig()
	fmt.Println("loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("config log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":9800")
}
