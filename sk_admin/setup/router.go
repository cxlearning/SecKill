package setup

import (
	"github.com/gin-gonic/gin"
	"log"
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
	router.GET("test", func(context *gin.Context) {

		context.Writer.Write([]byte("success"))
	})
}
