package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config       clientv3.Config
		client       *clientv3.Client
		err          error
		kv           clientv3.KV
		putOp, getOp clientv3.Op
		opResp       clientv3.OpResponse
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

	putOp = clientv3.OpPut("/cron/job3/job8", "job8")
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision：", opResp.Put().Header.Revision)

	getOp = clientv3.OpGet("/cron/job3/job8")
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据修改版本：", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("数据Value：", string(opResp.Get().Kvs[0].Value))
}
