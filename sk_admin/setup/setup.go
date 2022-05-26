package setup

import (
	"log"
	"myProject/SecKill/sk_admin/conf"
	"myProject/SecKill/sk_admin/library/etcd"
	"myProject/SecKill/sk_admin/library/logger"
)

func Run(configPath string) {

	logger.Init()
	conf.Init(configPath)
	etcd.Init()

}