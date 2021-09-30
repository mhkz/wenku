package router

import (
	"github.com/gin-gonic/gin"
	"wenku/controller/api"
)

func InitUserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("regist", api.Regist)
	}
}
