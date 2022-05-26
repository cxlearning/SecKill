package etcd

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var conn *clientv3.Client

func GetEtcdInstance() *clientv3.Client {
	return conn
}

func Init() {
	conn, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.20.10.18:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("connect to etcd failed, err:%v", err))
	}
	_ = conn
	fmt.Println("connect to etcd success")
}
