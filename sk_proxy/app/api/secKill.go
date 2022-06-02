package api

import (
	"github.com/gin-gonic/gin"
	"myProject/SecKill/sk_admin/app/response"
	"myProject/SecKill/sk_proxy/app/service"
	"myProject/SecKill/sk_proxy/memory"
	"time"
)

func SecKill(c *gin.Context) {

	var req memory.SecRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ClientErr(c, err.Error())
		return
	}

	req.ResultChan = make(chan *memory.SecResponse)
	req.CloseNotify = c.Writer.CloseNotify()
	req.AccessTime = time.Now().Unix()

	rep, err := service.SecKill(&req)
	if err != nil {
		response.Response(c, 0, err.Error(), "")
		return
	}
	response.Response(c, 0, "success", rep)
	return
}
