package router

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/handler"
)

type Router struct {
	helloHandler *handler.HelloHandler
}

func New(helloHandler *handler.HelloHandler) *Router {
	return &Router{
		helloHandler: helloHandler,
	}
}

func (r *Router) Register(rootGroup *gin.RouterGroup) {
	// swagger
	r.registerSwaggerRouter(rootGroup)
	//hello
	r.registerHelloRouter(rootGroup)

}
