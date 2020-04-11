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
		getResponse *clientv3.GetResponse
		watcher     clientv3.Watcher
		watchChan   clientv3.WatchChan
		err         error
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
	if getResponse, err = kv.Get(context.TODO(), "anhui"); err != nil {
		fmt.Println(err)
	}
	startReversion := getResponse.Header.Revision + 1
	//	watch
	watcher = clientv3.NewWatcher(client)
	watchChan = watcher.Watch(context.TODO(), "anhui", clientv3.WithRev(startReversion), clientv3.WithPrevKV())

	//通过for 遍历channel 只有channerl关闭程序才会退出
	for wat := range watchChan {
		for _, e := range wat.Events {
			switch int(e.Type) {
			case 0:
				if e.PrevKv != nil {
					fmt.Println(string(e.PrevKv.Value), "修改为", string(e.Kv.Value), e.Kv.CreateRevision, e.Kv.ModRevision)
				} else {
					fmt.Println("新写入", string(e.Kv.Value), e.Kv.CreateRevision, e.Kv.ModRevision)
				}

			case 1:
				fmt.Println("删除了", string(e.PrevKv.Value))
			}
		}
	}

}
