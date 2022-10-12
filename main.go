package main

import (
	"fmt"

	routes "C-SquaredService/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	r := gin.Default()               //1. 註冊一個路由器
	r.RedirectFixedPath = true       //   自動修正url 允許大小寫
	routes.NewController(r).Router() //2. 建立新的Router
	r.Run("0.0.0.0:8080")            //3. 執行（預設是23001埠）

	fmt.Println("end")
}
