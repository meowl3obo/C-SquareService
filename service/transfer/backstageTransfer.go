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
	productData.ParentClassify, err = strconv.Atoi(c.PostForm("parentClassify"))
	productData.ChildClassify, err = strconv.Atoi(c.PostForm("childClassify"))
	productData.Price, err = strconv.Atoi(c.PostForm("price"))
	productData.Intro = c.PostForm("intro")
	productData.Illustrate = c.PostForm("illustrate")
	productData.Name = c.PostForm("name")
	productData.Id = fmt.Sprintf("CSQ%v%v%v", productData.ParentClassify, productData.ChildClassify, hashService.Encrypt(productData.Name, 5))

	return productData, err
}
