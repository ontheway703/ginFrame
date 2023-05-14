package router

import (
	"game.interact.com/gin/frame/docs"
	"game.interact.com/gin/frame/handler/user"
	"game.interact.com/gin/frame/middleware"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	basePath = "/api"
	userBasePath = "/api/user"
)

var (
	userHandler user.Handler
)

func InitHandler() {
	userHandler = user.NewHandler()
}

func InitDocs()  {
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Title = "接口文档"
	docs.SwaggerInfo.Description = "小游戏接口文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "game.interact.io"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

}

func Register(r *gin.Engine) {
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(
		middleware.CommonMiddleware(),
	)
	UserRegister(r.Group(userBasePath))
}



