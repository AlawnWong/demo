package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		putResp *clientv3.PutResponse
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

	kv = clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "value1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("put Value = ", string(putResp.PrevKv.Value))
		}
	}
}