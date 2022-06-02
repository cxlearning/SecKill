package service

import (
	"encoding/json"
	"log"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/library/redis"
	"myProject/SecKill/sk_layer/memory"
	"time"
)

// 从redis读req 放到channel
func HandleReader() {

	log.Println("reader begin")

	conn := redis.GetInstance()

	for {
		//从Redis队列中取出数据
		data, err := conn.BRPop(time.Minute, conf.Config.Redis.Proxy2layerQueueName).Result()
		if err != nil {
			log.Printf("blpop from data failed, err : %v", err)
			continue
		}

		log.Printf("brpop from proxy to layer queue, data : %s\n", data)
		var req memory.SecRequest
		err = json.Unmarshal([]byte(data[1]), &req) // data[0]是listname
		if err != nil {
			log.Printf("blpop from data failed, err : %v", err)
			continue
		}

		//判断是否超时
		if time.Now().Unix()-req.AccessTime >= conf.Config.Server.MaxRequestWaitTimeout {
			log.Printf("req[%v] is expire", req)
			continue
		}

		//放入通道
		time := time.Tick(time.Duration(conf.Config.Server.SendToHandleChanTimeout) * time.Second)

		select {
		case <-time:
			log.Printf("send to handle chan timeout, req : %v", req)
		case memory.Mem.Read2HandleChan <- &req:
			break
		}
	}
}

// 从通道里读res 写到redis
func HandleWrite() {
	log.Println("HandleWrite, begin")

	for _res := range memory.Mem.Handle2WriteChan {
		err := sendToRedis(_res)
		if err != nil {
			log.Printf("send to redis, err : %v, res : %v", err, _res)
			continue
		}
		log.Println("send to redis success")
	}

}

//将数据推入到Redis队列
func sendToRedis(res *memory.SecResponse) (err error) {

	conn := redis.GetInstance()

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("json marshal err:", err)
		return
	}

	err = conn.LPush(conf.Config.Redis.Layer2proxyQueueName, string(data)).Err()
	if err != nil {
		log.Printf("rpush layer to proxy redis queue failed, err : %v", err)
		return
	}
	return nil
}
