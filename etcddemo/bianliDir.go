package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
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

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job0", "value0"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(putResp.Header.Revision)
	}

	// 使用 withPrefix 遍历目录
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)
	}
}

