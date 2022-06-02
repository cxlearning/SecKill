package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"myProject/SecKill/sk_layer/conf"
)

var conn *redis.Client

func GetInstance() *redis.Client {
	return conn
}
func Init() {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Config.Redis.HostPort,
		Password: conf.Config.Redis.Password,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Connect redis failed. Error : %v", err))
	}
	conn = client
	log.Println("redis connect success")
}

