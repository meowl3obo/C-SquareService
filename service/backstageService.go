package service

import (
	"fmt"
	"net/http"
	"os"

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
	}
	err, childClassifyList := provider.GetChildClassify()
	if err != nil {
		fmt.Println(fmt.Sprintf("Get ChildClassify, %v", err))
		c.JSON(http.StatusOK, classifyList)
	}
	classifyList = transfer.MergeClassify(parentClassifyList, childClassifyList)
	c.JSON(http.StatusOK, classifyList)
}

func InsertProduct(c *gin.Context) {
	productData, err := transfer.ProductFormToModel(c)
	if err != nil {
		c.JSON(http.StatusOK, ApiResponse{ResultCode: "500", ResultMessage: err})
		return
	}
	os.Mkdir("img", os.ModePerm)
	os.Mkdir(fmt.Sprintf("img/%v", productData.Id), os.ModePerm)
	mainImg, err := c.FormFile("mainImg")
	mainImgUrl := fmt.Sprintf("./img/%v/%v", productData.Id, mainImg.Filename)
	err = c.SaveUploadedFile(mainImg, mainImgUrl)
	if err != nil {
		c.JSON(http.StatusOK, ApiResponse{ResultCode: "500", ResultMessage: err})
		return
	}
	productData.MainImg = mainImgUrl

	c.JSON(http.StatusOK, productData)
}
