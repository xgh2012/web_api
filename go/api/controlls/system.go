package controlls

import (
	"demo/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type System struct {
}

func (system *System) Router(router *gin.Engine) {
	r := router.Group("/api/system/")
	r.GET("menu", system.menu)
}

func (system *System) menu(ctx *gin.Context) {
	menus := new(models.SystemMenu).GetAll()
	fmt.Printf("%+v", menus)

	result := make(map[string]interface{})
	result["code"] = 0
	result["count"] = 19
	result["data"] = menus
	result["msg"] = ""

	ctx.JSON(http.StatusOK, result)
}
