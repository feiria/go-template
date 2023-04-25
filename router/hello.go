package router

import (
	"github.com/gin-gonic/gin"
	"go-template/middleware"
)

func (r *Router) registerHelloRouter(Router *gin.RouterGroup) {
	helloRouterAllowCors := Router.Group("hello").Use(middleware.Cors())
	{
		helloRouterAllowCors.GET("/hello", r.helloHandler.Hello)
	}

	helloRouter := Router.Group("hello")
	{
		helloRouter.GET("/token", r.helloHandler.HelloGetToken)
	}

	helloRouterNeedToken := Router.Group("hello").Use(middleware.JWTAuth())
	{
		helloRouterNeedToken.POST("/auth", r.helloHandler.HelloNeedAuth)
	}

}
