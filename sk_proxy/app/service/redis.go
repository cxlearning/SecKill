package service

import (
	"encoding/json"
	"fmt"
	"log"
	"myProject/SecKill/sk_proxy/conf"
	"myProject/SecKill/sk_proxy/library/redis"
	"myProject/SecKill/sk_proxy/memory"
	"time"
)

/**
监听通道， 将req放到redis队列中
*/
func WriteHandle() {

	log.Println("req to redis list")
	for _req := range memory.Mem.SecReqChan {

		reqByte, err := json.Marshal(_req)
		if err != nil {
			log.Printf("json.Marshal req failed. Error : %v, req : %v", err, string(reqByte))
			continue
		}

		err = redis.GetInstance().LPush(conf.Config.Redis.Proxy2layerQueueName, string(reqByte)).Err()
		if err != nil {
			log.Printf("lpush req failed. Error : %v, req : %v", err, string(reqByte))
			continue
		}
		log.Printf("lpush req success. req : %v", string(reqByte))
	}

}

//从redis读取response, 并放到对应的接受通道
func ReadHandle() {

	conn := redis.GetInstance()

	for {
		data, err := conn.BRPop(time.Minute, conf.Config.Redis.Layer2proxyQueueName).Result()
		if err != nil {
			log.Printf("BRPop response failed. Error : %v, response : %v", err, data)
			continue
		}

		var res memory.SecResponse
		err = json.Unmarshal([]byte(data[1]), &res) //data[0]为list名
		if err != nil {
			log.Printf("json.Unmarshal failed. Error : %v", err)
			continue
		}
		log.Printf("rec response :%v", res)

		memory.Mem.UserConns.Lock.Lock()
		userKey := fmt.Sprintf("%s_%s", res.UserId, res.ProductId)
		memory.Mem.UserConns.Lock.Unlock()

		ch, ok := memory.Mem.UserConns.UserConnMap[userKey]
		if !ok {
			log.Printf("user not found : %v", userKey)
			continue
		}

		ch <- &res
		log.Printf("response send to chan succeee, userKey : %v", userKey)
	}

}
