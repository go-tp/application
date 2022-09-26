package main

import (
	"gtp/route"
)

func main() {
	
	// 路由
	r := route.Init()

	// 端口
	r.Run(":7777")
}