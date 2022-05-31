package service

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/memory"
	"myProject/SecKill/sk_layer/service/srv_err"
	"time"
)

/**
从通道读req处理，将res放到通道
*/
func HandleUser() {

	for _req := range memory.Mem.Read2HandleChan {
		res, err := handleSkill(_req)
		if err != nil {
			log.Printf("process request %v failed, err : %v", err)
			res = &memory.SecResponse{
				Code: srv_err.ErrServiceBusy,
			}
		}

		tick := time.Tick(time.Duration(conf.Config.Server.SendToWriteChanTimeout))

		select {
		case <-tick:
			log.Printf("send response timeout, res = %v\n", res)
			break
		case memory.Mem.Handle2WriteChan <- res:
		}
	}

}

/**
处理请求
*/
func handleSkill(req *memory.SecRequest) (res *memory.SecResponse, err error) {

	//加锁获取内存中商品数据
	memory.Mem.Products.Lock.RLock()
	defer memory.Mem.Products.Lock.RUnlock()

	res.ProductId = req.ProductId
	res.UserId = req.UserId

	product, ok := memory.Mem.Products.ProductMap[req.ProductId]
	if !ok {
		log.Printf("not found product : %v", req.ProductId)
		res.Code = srv_err.ErrNotFoundProduct
		return
	}

	if product.Status == srv_err.ProductStatusSoldout {
		res.Code = srv_err.ErrSoldout
		return
	}

	// todo 每秒卖出数量限制， 单人购买限制

	nowTime := time.Now().Unix()

	// 是否有货
	count := memory.Mem.ProductSoldMgr.Count(req.ProductId)
	if count >= product.Total {
		res.Code = srv_err.ErrSoldout
		return
	}

	if rand.Float64() < product.BuyRate {
		res.Code = srv_err.ErrRetry
		return
	}

	//可以买
	memory.Mem.ProductSoldMgr.Add(req.ProductId, 1)

	//用户Id、商品id、当前时间、密钥
	res.Code = srv_err.ErrSecKillSucc
	tokenData := fmt.Sprintf("userId=%d&productId=%d&timestamp=%d&security=%s", req.UserId, req.ProductId, nowTime, conf.Config.Server.TokenPassWd)
	res.Token = fmt.Sprintf("%x", md5.Sum([]byte(tokenData))) //MD5加密
	res.TokenTime = nowTime

	return

}
