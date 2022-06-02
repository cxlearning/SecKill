package setup

import (
	"myProject/SecKill/sk_proxy/app/service"
	"myProject/SecKill/sk_proxy/conf"
	"myProject/SecKill/sk_proxy/library/etcd"
	"myProject/SecKill/sk_proxy/library/logger"
	"myProject/SecKill/sk_proxy/library/redis"
)

func Run(configPath string) {

	logger.Init()
	conf.Init(configPath)
	etcd.Init()
	redis.Init()

	service.Init()
	initServer(conf.Config.Server.Port)
}
