package main

import (
	routes "C-SquaredService/controller"
	"C-SquaredService/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //1. 註冊一個路由器
	r.Use(cors.Default())
	r.Use(service.CorsMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour
	corsConfig.AllowOriginFunc = func(origin string) bool {
		return true
	}
	r.RedirectFixedPath = true       //   自動修正url 允許大小寫
	routes.NewController(r).Router() //2. 建立新的Router
	r.Run("0.0.0.0:8080")            //3. 執行（預設是23001埠）
}
