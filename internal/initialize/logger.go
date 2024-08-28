package initialize

import (
	"github.com/go-ecommerce/global"
	"github.com/go-ecommerce/pkg/logger"
)

func InitLoger() {
	global.Logger = logger.NewLogger()
}