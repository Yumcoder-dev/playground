// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//

// test connection to etcd cluster
package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

// connect to yumd-etcd1, yumd-etcd2, and yumd-etcd3
// docker exec -it yumd-etcd1 sh
// export ETCDCTL_API=3
// etcdctl get sample_key

func getCli() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.20.10.10:2379", "10.20.10.11:2379", "10.20.10.12:2379"},
		DialTimeout: 5 * time.Second,
	})
}

// run te test when docker yumd-etcd1(10.20.10.10) is stopped
// driver automatically switch to yumd-etcd2(10.20.10.11)
func Test_Cluster(t *testing.T) {
	cli, err := getCli()
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			t.Errorf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			t.Errorf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			t.Errorf("client-side error: %v", err)
		default:
			t.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	t.Log(resp)
}

// for test grafana, see in the panel http://localhost:3000
// n : gRpc ops
// n/2: put
// n/2: get/range
func Test_Put_GET(t *testing.T) {
	n := 400
	cli, err := getCli()
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer cli.Close()

	for i := 0; i < n; i++ {
		ctx, _ := context.WithCancel(context.Background())
		if i%2 == 0 {
			cli.Get(ctx, fmt.Sprintf("/%d/sample_key", i))
		} else {
			cli.Put(ctx, fmt.Sprintf("/%d/sample_key", i), "val")
		}
	}
}

// for test grafana, see in the panel http://localhost:3000
// 1: watch
func Test_Watcher(t *testing.T) {
	cli, err := getCli()
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer cli.Close()

	rch := cli.Watch(context.Background(), "foo")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
