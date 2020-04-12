package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		client  *clientv3.Client
		kv      clientv3.KV
		PutOpt  clientv3.Op
		putResp clientv3.OpResponse
		getResp clientv3.OpResponse
		err     error
	)
	config := clientv3.Config{
		Endpoints:   []string{"47.244.184.242:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)
	PutOpt = clientv3.OpPut("anhui", "123123")
	if putResp, err = kv.Do(context.TODO(), PutOpt); err != nil {
		fmt.Println(err)
	}
	fmt.Println("写入数据", putResp.Put().Header.Revision)

	Getop := clientv3.OpGet("anhui")
	if getResp, err = kv.Do(context.TODO(), Getop); err != nil {
		fmt.Println(err)
	}
	fmt.Println(getResp.Get().Kvs[0].ModRevision)
	fmt.Println(string(getResp.Get().Kvs[0].Value))
}
