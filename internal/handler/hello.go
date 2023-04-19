package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-template/global"
	"go-template/internal/logic"
	"go-template/internal/types/request"
	"go-template/internal/types/response"
	"go-template/middleware"
	"go.uber.org/zap"
)

type HelloHandler struct {
	helloLogic logic.HelloLogic
}

func NewHelloHandler(helloLogic logic.HelloLogic) *HelloHandler {
	return &HelloHandler{
		helloLogic: helloLogic,
	}
}

func (h *HelloHandler) Hello(c *gin.Context) {
	resp, err := h.helloLogic.HelloAPI()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.SuccessWithData(resp, c)
	return
}

func (h *HelloHandler) HelloNeedAuth(c *gin.Context) {
	resp, err := h.helloLogic.HelloAPI()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.SuccessWithData(resp, c)
	return
}

func (h *HelloHandler) HelloGetToken(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	TokenNext(request.BaseClaims{
		ID:   26,
		Uuid: "5d586a1f-3a02-4fca-860e-4f3a1de5c063",
		Name: "gailun2",
	}, token, c)
	return
}

func TokenNext(baseClaims request.BaseClaims, t string, c *gin.Context) {
	j := &middleware.JWT{SigningKey: []byte(viper.GetString("jwt.signing-key"))}
	claims := j.CreateClaims(baseClaims)

	var token string
	var err error
	if t == "" {
		token, err = j.CreateToken(claims)
	} else {
		global.GetLog().Info("refresh token")
		token, err = j.RefreshToken(t, claims)
	}

	if err != nil {
		global.GetLog().Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.SuccessWithDetail("login success", struct {
		Token     string
		ExpiresAt int64
	}{
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}, c)
	return
}
