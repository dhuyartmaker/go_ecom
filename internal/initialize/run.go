package initialize

import (
	"fmt"

	"github.com/go-ecommerce/global"
)

func Initialize() {
	LoadConfig()
	fmt.Print("Load config:", global.Config.Server.Port)
	InitLoger()
	global.Logger.Info("Config Logger")
	InitMysql()
	InitRedis()
	InitRouter()

}
