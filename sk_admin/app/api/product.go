package api

import (
	"github.com/gin-gonic/gin"
	"myProject/SecKill/sk_admin/app/model"
	"myProject/SecKill/sk_admin/app/response"
	"myProject/SecKill/sk_admin/app/service"
)

var ProductApi productApi

type productApi struct {
}

func (*productApi) Creat(c *gin.Context) {

	var req model.ProductCreatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ClientErr(c, err.Error())
		return
	}

	if err := service.ProductService.Creat(req); err != nil {
		response.ServerErr(c, err.Error())
		return
	}

	response.Success(c)
	return
}
