package controller

import (
	"fmt"
	"net/http"

	"C-SquaredService/service"

	"github.com/gin-gonic/gin"
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
			v1.Use(service.DBMiddleware())
			v1.GET("/param/:first", Param)
			v1.GET("/query", service.Test)
			v1.POST("/user", service.CreateUser)
		}
		back := api.Group("/back")
		{
			back.Use(service.DBMiddleware())
			back.GET("/classify", service.GetProductClassify)
			back.POST("/product", service.InsertProduct)
		}
	}
}
func version(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%v V:%v-%v,Build:%v", "demo", "1.0.1", "Local", "2022/01/01 15:32:10"))
}

func Param(c *gin.Context) {
	firstParameter := c.Param("first")
	c.String(http.StatusOK, "Param = %v\n", firstParameter)
}
