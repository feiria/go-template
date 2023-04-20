package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) registerSwaggerRouter(rootGroup *gin.RouterGroup) {
	swagHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	if swagHandler != nil {
		swaggerRouter := rootGroup.Group("swagger")
		{
			swaggerRouter.GET("/*any", swagHandler)
		}
	}
}
