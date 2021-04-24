package api

import (
	"demo/api/adapts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
}

func (l *Index) Router(router *gin.Engine) {
	r := router.Group("/api/index/")

	r.GET("init", l.init)
}

//获取菜单配置等信息 ---------------- 开始
type HomeInfoStr struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}
type LogoInfoStr struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Href  string `json:"href"`
}
type MenuInfoStr struct {
	Title  string        `json:"title"`
	Icon   string        `json:"icon"`
	Href   string        `json:"href"`
	Target string        `json:"target"`
	Child  []MenuInfoStr `json:"child"`
}

type BarInfoStr struct {
	BarId      string `json:"bar_id"`
	BarName    string `json:"bar_name"`
	Permission int    `json:"permission"`
	Version    string `json:"version"`
}

type Menu struct {
	HomeInfo HomeInfoStr   `json:"homeInfo"`
	LogoInfo LogoInfoStr   `json:"logoInfo"`
	MenuInfo []MenuInfoStr `json:"menuInfo"`
	BarInfo  BarInfoStr    `json:"barInfo"`
}

//获取验证码
func (l *Index) init(ctx *gin.Context) {
	systemInit := new(adapts.SystemInit).GetSystemInit()
	ctx.JSON(http.StatusOK, systemInit)
}
