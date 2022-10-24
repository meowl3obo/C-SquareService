package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	. "C-SquaredService/model"
	provider "C-SquaredService/provider"
	transfer "C-SquaredService/service/transfer"
)

func GetProductClassify(c *gin.Context) {
	var classifyList = []Classify{}
	err, parentClassifyList := provider.GetParentClassify()
	if err != nil {
		fmt.Println(fmt.Sprintf("Get ParentClassify, %v", err))
		c.JSON(http.StatusOK, classifyList)
		return
	}
	err, childClassifyList := provider.GetChildClassify()
	if err != nil {
		fmt.Println(fmt.Sprintf("Get ChildClassify, %v", err))
		c.JSON(http.StatusOK, classifyList)
		return
	}
	classifyList = transfer.MergeClassify(parentClassifyList, childClassifyList)
	c.JSON(http.StatusOK, classifyList)
}

func InsertProduct(c *gin.Context) {
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
