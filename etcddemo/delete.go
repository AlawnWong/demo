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
		delResp *clientv3.DeleteResponse
		putResp *clientv3.PutResponse
		kv      clientv3.KV
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

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job0", "job0"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("put job0 Revision = ", putResp.Header.Revision)
	}

	// 如果要删除多个key，可以考虑使用 clientv3.WithPrefix() 选项
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job0", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}

	if len(delResp.PrevKvs) != 0 {
		for _, keyPair := range delResp.PrevKvs {
			fmt.Println(string(keyPair.Value), keyPair.CreateRevision, keyPair.ModRevision)
		}
	}
}
