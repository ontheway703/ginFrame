package router

import (
	"game.interact.com/gin/frame/handler/user"
	"game.interact.com/gin/frame/middleware"
	"github.com/gin-gonic/gin"
)

const (
	userBasePath = "user"
)

var (
	userHandler user.Handler
)

func InitHandler() {
	userHandler = user.NewHandler()
}

func Register(r *gin.Engine) {
	r.Use(
		middleware.CommonMiddleware(),
	)
	UserRegister(r.Group(userBasePath))
}
