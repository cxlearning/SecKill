package etcd

import (
	"fmt"
	"myProject/SecKill/sk_proxy/conf"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var conn *clientv3.Client

func GetEtcdInstance() *clientv3.Client {
	return conn
}

func Init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Config.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("connect to etcd failed, err:%v", err))
	}
	conn = cli
}
