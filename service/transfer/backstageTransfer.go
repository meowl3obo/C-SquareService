package transfer

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	hashService "C-SquaredService/hashService"
	. "C-SquaredService/model"
)

func ProductFormToModel(c *gin.Context) (Product, error) {
	var productData = Product{}
	err := errors.New("")
	productData.Status, err = strconv.Atoi(c.PostForm("status"))
	productData.ParentClassify = c.PostForm("parentClassify")
	productData.ChildClassify = c.PostForm("childClassify")
	productData.Price, err = strconv.Atoi(c.PostForm("price"))
	productData.Intro = c.PostForm("intro")
	productData.Illustrate = c.PostForm("illustrate")
	productData.Name = c.PostForm("name")
	productData.Id = fmt.Sprintf("CSQ%v%v%v", productData.ParentClassify, productData.ChildClassify, hashService.Encrypt(productData.Name, 5))

	return productData, err
}

func InventoryFormToModel(c *gin.Context, productID string) (Inventory, error) {
	var InventoryData = Inventory{}
	err := errors.New("")
	InventoryData.ProductId = productID
	InventoryData.Color = c.PostForm("color")
	InventoryData.Size, err = strconv.ParseFloat(c.PostForm("size"), 64)
	InventoryData.Unit = c.PostForm("unit")
	InventoryData.PreOrderAmount, err = strconv.Atoi(c.PostForm("preOrderAmount"))
	InventoryData.NowAmount, err = strconv.Atoi(c.PostForm("nowAmount"))
	InventoryData.Id = fmt.Sprintf("CSQI%v", hashService.Encrypt(fmt.Sprintf("%v%v%v%v", InventoryData.ProductId, InventoryData.Color, InventoryData.Size, InventoryData.Unit), 6))

	return InventoryData, err
}
