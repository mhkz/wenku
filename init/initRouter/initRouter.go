package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "wenku/docs"
	"wenku/router"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Router.Use(middleware.Logger())
	router.InitUserRouter(Router)
	return Router
}
