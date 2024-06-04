package main

import (
	"go-ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()
	err := r.Run(":9800")
	if err != nil {
		return
	}
}
