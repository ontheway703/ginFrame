package main

import (
	"game.interact.com/gin/frame/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// client 初始化
func init() {
}

// main
func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"now":     time.Now(),
		})
	})

	defer r.Run()

	router.InitHandler()
	router.Register(r)
}
