package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"myProject/SecKill/sk_layer/conf"
)

var conn redis.Conn

func GetInstance() redis.Conn {
	return conn
}
func Init() {
	var err error
	conn, err = redis.Dial("tcp", conf.Config.Redis.HostPort)
	if err != nil {
		panic(fmt.Sprintf("conn redis failed, err = %s", err.Error()))
	}
}
