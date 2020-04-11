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
		getResponse *clientv3.GetResponse
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

	// Put
	if putResponse, err = Kv.Put(context.TODO(), "aa", "hello", clientv3.WithPrevKV()); err != nil {
		fmt.Println("put error:", err)
	}

	fmt.Println("Revision:", putResponse.Header.Revision)
	if putResponse.PrevKv != nil {
		fmt.Println("prekv:", string(putResponse.PrevKv.Value))
	}

	// Get
	if getResponse, err = Kv.Get(context.TODO(), "aa"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(getResponse.Kvs, getResponse.Count)

	//[key:"aa" create_revision:28 mod_revision:28 version:1 value:"hello" ] 1
	//这里的 create_revision mod_revision 是etcd 总的操作次数

	fmt.Println("success")
}
