package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "C-SquaredService/model"
	provider "C-SquaredService/provider"
)

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func Test(c *gin.Context) {
	secondParameter := c.Query("second")
	response := provider.GetTest(secondParameter)
	userData := Users{}
	provider.CreateUser(userData)
	c.Data(http.StatusOK, "application/json", response)
}
