package service

import (
	"errors"
	"fmt"
	"myProject/SecKill/sk_proxy/conf"
	"myProject/SecKill/sk_proxy/memory"
	"time"
)

func SecKill(req *memory.SecRequest) (memory.SecResponse, error) {

	//拿内存中的商品数据， 加读锁
	memory.Mem.Products.Lock.RLock()
	defer memory.Mem.Products.Lock.RUnlock()

	_, ok  := memory.Mem.Products.ProductMap[req.ProductId]
	if !ok {
		return memory.SecResponse{}, errors.New(fmt.Sprintf("productID:%s not exist", req.ProductId))
	}

	tick := time.Tick(time.Duration(conf.Config.Server.SendReq2ChanTimeOut) * time.Second)

	userKey := fmt.Sprintf("%s_%s", req.UserId, req.ProductId)
	memory.Mem.UserConns.Lock.Lock()
	memory.Mem.UserConns.UserConnMap[userKey] = req.ResultChan
	memory.Mem.UserConns.Lock.Unlock()
	defer func() {
		memory.Mem.UserConns.Lock.Lock()
		delete(memory.Mem.UserConns.UserConnMap, userKey)
		memory.Mem.UserConns.Lock.Unlock()
	}()
	// 发送到通道
	select {
	case <-tick:
		return memory.SecResponse{}, errors.New("send req to chan timeout")
	case memory.Mem.SecReqChan <- req:
	}

	tickW := time.Tick(time.Duration(conf.Config.Server.WaitResponseTimeOut)* time.Second)

	select {
	case rep := <-req.ResultChan:
		return *rep, nil
	case <- req.CloseNotify:
		return memory.SecResponse{}, errors.New("客户端结果关闭")
	case <-tickW:
		return memory.SecResponse{}, errors.New("等待layer层结果超时")
	}
}
