package main

import (
	"game.interact.com/gin/frame/router"
	"github.com/gin-gonic/gin"
)

// client 初始化
func init() {
}

// main
func main() {
	r := gin.Default()
	defer r.Run()

	router.InitDocs()
	router.InitHandler()
	router.Register(r)
}
