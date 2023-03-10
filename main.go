package main

import (
	"ltt-gc/config"
	"ltt-gc/router"
)

func main() {
	config.Init()
	r := router.NewRouter()
	r.Run(":8899")
}
