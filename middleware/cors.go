package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	//todo 可在配置文件读取允许跨域列表
	allowOrigins = []string{
		"http://127.0.0.1:8867",
		"https://127.0.0.1:8867",
		"http://127.0.0.1:8868",
		"https://127.0.0.1:8868",
	}

	allowMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,

		http.MethodPatch, // RFC 5789
	}

	allowHeaders = []string{
		"Content-Type",
		"Access-Token",
		"Authorization",
	}

	exposeHeaders = []string{
		"Content-Length",
	}
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		// AllowAllOrigins: true,
		// 接受哪些域名的请求
		AllowOrigins:  allowOrigins,
		AllowMethods:  allowMethods,
		AllowHeaders:  allowHeaders,
		ExposeHeaders: exposeHeaders,
		// 是否允许发送Cookie
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		// default 12h
		MaxAge: 6 * time.Hour,
	}

	//开启debug则默认不阻止任何跨域请求
	if viper.GetBool("server.debug") {
		c.AllowAllOrigins = true
	} else {
		c.AllowOrigins = allowOrigins
	}

	return cors.New(c)
}
