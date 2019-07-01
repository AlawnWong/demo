package main


import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchRespChan      clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		ctx                context.Context
		cancelFunc         context.CancelFunc
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
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")
			kv.Delete(context.TODO(), "/cron/jobs/job7")
			time.Sleep(1 * time.Second)
		}
	}()

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值", string(getResp.Kvs[0].Value))
	}

	watchStartRevision = getResp.Header.Revision + 1

	watcher = clientv3.NewWatcher(client)
	fmt.Println("从该版本向后监听", watchStartRevision)

	ctx, cancelFunc = context.WithCancel(context.TODO())
	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	for watchResp = range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision：", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", event.Kv.ModRevision)
			}
		}
	}
}