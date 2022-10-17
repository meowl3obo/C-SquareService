package control

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	service "C-SquaredService/service"
)

type Controller struct {
	*gin.Engine
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

// 設定Router
func (r *Controller) Router() {
	api := r.Group("/api")
	{
		api.GET("/version", version)
		v1 := api.Group("/v1")
		{
			v1.GET("/param/:first", Param)
			v1.GET("/query", service.Test)
			v1.POST("/user", service.CreateUser)
		}
	}
}

// @Summary 版本查詢
// @Tags Version
// @version 1.0
// @Router /version [get]
func version(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%v V:%v-%v,Build:%v", "demo", "1.0.1", "Local", "2022/01/01 15:32:10"))
}

// @Summary Param 範例
// @Tags Param
// @version 1.0
// @Router /param/test [get]
func Param(c *gin.Context) {
	firstParameter := c.Param("first")
	c.String(http.StatusOK, "Param = %v\n", firstParameter)
}
