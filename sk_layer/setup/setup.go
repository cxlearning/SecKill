package setup

import (
	"myProject/SecKill/sk_layer/app/memory"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/library/etcd"
	"myProject/SecKill/sk_layer/library/logger"
)

func Run(configPath string) {

	logger.Init()
	conf.Init(configPath)
	etcd.Init()

	memory.Init()

}