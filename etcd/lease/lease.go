package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		client        *clientv3.Client
		getResponse   *clientv3.GetResponse
		lease         clientv3.Lease
		leaseResponse *clientv3.LeaseGrantResponse
		LeaseChan     <-chan *clientv3.LeaseKeepAliveResponse
		leasechans    *clientv3.LeaseKeepAliveResponse
		err           error
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
	// lease 对象
	lease = clientv3.NewLease(client)
	leaseResponse, err = lease.Grant(context.TODO(), 10)
	lease_id := leaseResponse.ID

	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	LeaseChan, err = lease.KeepAlive(ctx, lease_id)

	go func() {
		for {
			select {
			case leasechans = <-LeaseChan:
				if leasechans != nil {
					fmt.Println("续约成功:", leasechans.ID)
				} else {
					fmt.Println("租期失效了")
					// golang goto使用
					goto END
				}
			}
		}
	END:
	}()

	// kv对象
	kv := clientv3.NewKV(client)
	kv.Put(context.TODO(), "anhui", "aa", clientv3.WithLease(lease_id))

	for {
		getResponse, err = kv.Get(context.TODO(), "anhui", clientv3.WithPrevKV())
		if getResponse.Count > 0 {
			fmt.Println("拿到了", getResponse.Kvs)
		} else {
			fmt.Println("过期了")
			break
		}
		time.Sleep(2 * time.Second)
	}

}
