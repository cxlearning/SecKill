package setup

import (
	"myProject/SecKill/sk_admin/conf"
	"myProject/SecKill/sk_admin/library/db"
	"myProject/SecKill/sk_admin/library/etcd"
	"myProject/SecKill/sk_admin/library/logger"
)

func Run(configPath string) {

	logger.Init()
	conf.Init(configPath)
	etcd.Init()
	db.Init()

	initServer(conf.Config.Server.Port)

}