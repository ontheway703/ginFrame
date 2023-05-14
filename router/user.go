package router

import (
	"game.interact.com/gin/frame/middleware"
	"game.interact.com/gin/frame/model"
	"game.interact.com/gin/frame/util/decorator"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {

	router.Use(middleware.CustomMiddleware())

	router.Any("/info", decorator.ApiDecorator(
		userHandler.Info,
		&model.UserInfoReq{},
	))
}
