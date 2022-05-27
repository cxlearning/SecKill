package api

import (
	"github.com/gin-gonic/gin"
	"myProject/SecKill/sk_admin/app/model"
	"myProject/SecKill/sk_admin/app/response"
	"myProject/SecKill/sk_admin/app/service"
)

var ActivityApi activity

type activity struct {
}

func (a *activity) Creat(c *gin.Context) {
	var req model.ActivityCreatReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ClientErr(c, err.Error())
		return
	}

	if err := service.ActivityService.Creat(req); err != nil {
		response.ServerErr(c, err.Error())
		return
	}

	response.Success(c)
	return
}
