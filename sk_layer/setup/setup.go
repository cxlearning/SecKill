package setup

import (
	"myProject/SecKill/sk_layer/app/memory"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/library/etcd"
	"myProject/SecKill/sk_layer/library/logger"
	"myProject/SecKill/sk_layer/library/redis"
)

func Run(configPath string) {

	logger.Init()
	conf.Init(configPath)
	etcd.Init()
	redis.Init()

	memory.Init()

}