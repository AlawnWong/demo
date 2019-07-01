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
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseId        clientv3.LeaseID
		leaseGrantResp *clientv3.LeaseGrantResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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
	if leaseGrantResp, err = lease.Grant(context.TODO(), 2); err != nil {
		fmt.Println(err)
		return
	}

	leaseId = leaseGrantResp.ID
	// 自动续租
	ctx, cancelFunc = context.WithCancel(context.TODO())
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)
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
	txn = kv.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9"))

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	// 是否抢到锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 处理任务
	fmt.Println("处理任务中")
	time.Sleep(10 * time.Second)
}
