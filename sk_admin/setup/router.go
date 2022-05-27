package setup

import (
	"github.com/gin-gonic/gin"
	"log"
	"myProject/SecKill/sk_admin/app/api"
)

//初始化Http服务
func initServer(host string) {
	router := gin.Default()
	setupRouter(router)
	err := router.Run(host)
	if err != nil {
		log.Printf("Init http server. Error : %v", err)
	}
}

func setupRouter(router *gin.Engine) {
 	router.POST("product/creat", api.ProductApi.Creat)
	router.POST("activity/creat", api.ActivityApi.Creat)
}
