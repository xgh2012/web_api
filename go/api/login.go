package api

import (
	"demo/utils"
	"github.com/gin-gonic/gin"
)

type Login struct {

}
func (l *Login) Router(router *gin.Engine) {
	r := router.Group("/api/login/")

	r.GET("captcha", l.captcha)
}

//获取验证码
func (l *Login) captcha(ctx *gin.Context) {
	utils.GetCaptcha(ctx)
}

