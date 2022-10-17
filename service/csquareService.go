package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	provider "C-SquaredService/provider"
)

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func Test(c *gin.Context) {
	secondParameter := c.Query("second")
	response := provider.GetTest(secondParameter)
	provider.CreateUser()
	c.Data(http.StatusOK, "application/json", response)
}
