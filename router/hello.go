package router

import (
	"github.com/gin-gonic/gin"
	"go-template/middleware"
)

func (r *Router) registerHelloRouter(Router *gin.RouterGroup) {
	helloRouter := Router.Group("hello")
	{
		helloRouter.GET("/token", r.helloHandler.HelloGetToken)
		helloRouter.GET("/hello", r.helloHandler.Hello)
	}

	helloRouterNeedToken := Router.Group("hello").Use(middleware.JWTAuth())
	{
		helloRouterNeedToken.POST("/auth", r.helloHandler.HelloNeedAuth)
	}

}
