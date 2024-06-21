package initialize

import (
	"fmt"
	"go-ecommerce/global"
)

// Run all initialize
func Run() {
	// load config
	LoadConfig()
	fmt.Println("loading configuration mysql", global.Config.Mysql.Username)
	InitLooger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":9800")
}
