package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		putResp        *clientv3.PutResponse
		leaseId        clientv3.LeaseID
		kv             clientv3.KV
		getResp        *clientv3.GetResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	lease = clientv3.NewLease(client)

	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	leaseId = leaseGrantResp.ID

	// 自动续租
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已尽失效了")
					goto END
				} else {
					fmt.Println("续租正常", keepResp.ID)
				}
			}
		}
	END:
	}()

	kv = clientv3.NewKV(client)

	// 使用 leaseid 创建
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "hello world", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功", putResp.Header.Revision)

	// 检查是否过期
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1", clientv3.WithCountOnly()); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("/cron/lock/job1过期了")
			break
		}
		time.Sleep(2 * time.Second)
	}
}

