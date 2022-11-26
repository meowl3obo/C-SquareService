package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "C-SquaredService/model"
	provider "C-SquaredService/provider"
	"C-SquaredService/service/transfer"
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

func GetProducts(c *gin.Context) {
	parentID := c.Param("id")
	childID := c.Query("child")
	if parentID == "sale" {
		err, products := provider.GetSaleProducts()
		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, products)
		}
	} else if parentID == "new" {
		err, products := provider.GetNewProducts()
		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, products)
		}
	} else {
		err, products := provider.GetProducts(parentID, childID)
		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, products)
		}
	}
}

func GetProduct(c *gin.Context) {
	ID := c.Param("id")
	err, product := provider.GetProduct(ID)
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		err, inventorys := provider.GetProductInventory(ID)
		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			productInventory := transfer.MergeProductInventory(product, inventorys)
			c.JSON(http.StatusOK, productInventory)
		}
	}
}
