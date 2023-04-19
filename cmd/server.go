package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-template/global"
	"go-template/pkg/config"
	"go-template/router"
	"net/http"
	"time"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
}

func (s *Server) Start() *http.Server {
	routerGroup := s.engine.Group("/api")
	s.apiRouter.Register(routerGroup)

	return &http.Server{
		Addr:         viper.GetString("server.addr"),
		Handler:      s.engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func New(apiRouter *router.Router) *Server {
	var ginMode string
	if viper.GetBool("server.log") {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	if ginMode != gin.ReleaseMode {
		handlerFunc := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s|%s|%d|%s\n",
					params.Method,
					params.Path,
					params.StatusCode,
					params.ClientIP,
				)
			},
			Output: &config.Output{Logger: global.Log},
		})
		engine.Use(handlerFunc)
	}

	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}
