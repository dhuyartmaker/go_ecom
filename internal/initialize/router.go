package initialize

import (
	"github.com/go-ecommerce/internal/router"
)

func InitRouter() {
	r := router.NewRouter();
	r.Run(":3000")
}