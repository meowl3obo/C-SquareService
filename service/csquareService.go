package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "C-SquaredService/model"
	provider "C-SquaredService/provider"
)

func Test(c *gin.Context) {
	secondParameter := c.Query("second")
	response := provider.GetTest(secondParameter)
	c.Data(http.StatusOK, "application/json", response)
}

func CreateUser(c *gin.Context) {
	userData := User{
		Email:    "dear91304526@gmail.com",
		Password: "meowl870706",
		Status:   0,
		CreateAt: "2022-10-17",
		UpdateAt: "2022-10-17",
		Name:     "安安",
	}
	err := provider.CreateUser(userData)
	response := ApiResponse{}
	if err != nil {
		response.ResultCode = "500"
		response.ResultMessage = err
	} else {
		response.ResultCode = "200"
		response.ResultMessage = "success"
	}
	c.JSON(http.StatusOK, response)
}
