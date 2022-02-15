package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/gink/controller/api"
)

func apiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.Use(throttle())
	apiRouter.Use(mustAuth())
	{
		apiRouter.POST("/user", format(api.DefaultUserController.GetUserInfo))
	}
}
