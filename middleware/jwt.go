package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go-template/internal/types/request"
	"go-template/internal/types/response"
	"go-template/pkg/times"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter jwt")
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage("token is null", c)
			c.Abort()
			return
		}
		fmt.Println(token)
		j := request.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, request.TokenExpired) {
				response.FailWithMessage("授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := times.ParseDuration(viper.GetString("jwt.buffer-time"))
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("x-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
