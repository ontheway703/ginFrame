package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CommonMiddleware")

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("CommonMiddleware err", err)
			}
		}()
		c.Next()
	}
}

func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CustomMiddleware")
		c.Next()
	}
}
