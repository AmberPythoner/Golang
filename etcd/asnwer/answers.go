package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		client *clientv3.Client
		//getResponse *clientv3.GetResponse
		//watcher     clientv3.Watcher
		//watchChan   clientv3.WatchChan
		err error
	)
	config := clientv3.Config{
		Endpoints:   []string{"47.244.184.242:2379"}, //集群列表
		DialTimeout: 5 * time.Second,
	}
	//client对象
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// kv对象
	kv := clientv3.NewKV(client)

	//获取KV数据
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	//1）在课程里newKV返回的对象，内部是无限重试的。
	//2）golang打断超时的机制是监听context.Done()，所以如果阻塞在没有监听context.Done()的方法上，是无法取消打断的。
	getResp, err := kv.Get(ctx, "/cron/bin/job2")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}
}
