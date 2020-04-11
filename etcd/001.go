package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		client      *clientv3.Client
		putResponse *clientv3.PutResponse
		err         error
	)
	config := clientv3.Config{
		Endpoints:   []string{"47.244.184.242:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	Kv := clientv3.NewKV(client)
	if putResponse, err = Kv.Put(context.TODO(), "name", "13", clientv3.WithPrevKV()); err != nil {
		fmt.Println("put error:", err)
	}

	fmt.Println("Revision:", putResponse.Header.Revision)
	fmt.Println("prekv:", string(putResponse.PrevKv.Value))

	fmt.Println("success")
}
