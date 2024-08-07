package main

import (
	"github.com/go-ecommerce/internal/router"
)

func main() {
	r := router.NewRouter();
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
