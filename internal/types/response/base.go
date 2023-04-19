package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type EmptyRetData struct{}

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "success", EmptyRetData{}, c)
}

func SuccessWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, msg, EmptyRetData{}, c)
}

func SuccessWithData(data any, c *gin.Context) {
	Result(SUCCESS, "success", data, c)
}

func SuccessWithDetail(msg string, data any, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "fail", EmptyRetData{}, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, msg, EmptyRetData{}, c)
}

func FailWithDetail(msg string, data any, c *gin.Context) {
	Result(ERROR, msg, data, c)
}
