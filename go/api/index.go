package api

import (
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
	init :=`{
  "homeInfo": {
    "title": "首页",
    "href": "page/welcome-1.html?t=1"
  },
  "logoInfo": {
    "title": "LAYUI MINI",
    "image": "images/logo.png",
    "href": ""
  },
  "menuInfo": [
    {
      "title": "常规管理",
      "icon": "fa fa-address-book",
      "href": "",
      "target": "_self",
      "child": [
        {
          "title": "系统设置",
          "href": "",
          "icon": "fa fa-home",
          "target": "_self",
          "child": [
            {
              "title": "用户管理",
              "href": "page/welcome-1.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },
            {
              "title": "角色管理",
              "href": "page/welcome-2.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },
            {
              "title": "权限管理",
              "href": "page/welcome-3.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },
            {
              "title": "菜单管理",
              "href": "page/welcome-3.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },
            {
              "title": "操作日志",
              "href": "page/welcome-3.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            }
          ]
        },
		{
          "title": "数据管理",
          "href": "",
          "icon": "fa fa-home",
          "target": "_self",
          "child": [
            {
              "title": "场所列表",
              "href": "page/data/list.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },
            {
              "title": "数据录入",
              "href": "page/data/add.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            },{
              "title": "数据审核",
              "href": "page/menu.html",
              "icon": "fa fa-tachometer",
              "target": "_self"
            }
          ]
        }
      ]
    }
  ]
}`

	ctx.String(http.StatusOK,init)
}

