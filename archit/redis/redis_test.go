// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package redis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

func getCli() (*redisc.Cluster, error) {
	cluster := &redisc.Cluster{
		StartupNodes: []string{
			"10.20.20.10:7000",
			"10.20.20.11:7000",
			"10.20.20.12:7000",
			"10.20.20.13:7000",
			"10.20.20.14:7000",
			"10.20.20.15:7000",
		},
		DialOptions: []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool:  createPool,
	}

	// initialize its mapping
	err := cluster.Refresh()
	return cluster, err
}

func createPool(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, opts...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

func Test_Cluster(t *testing.T) {
	n := 10000
	cluster, err := getCli()
	if err != nil {
		t.Fatalf("Init redis cluster failed: %v", err)
	}
	defer cluster.Close()

	// get a connection from the cluster
	conn := cluster.Get()
	defer conn.Close()

	// create the retry connection - only Do, Close and Err are
	// supported on that connection. It will make up to 3 attempts
	// to get a valid response, and will wait 100ms before a retry
	// in case of a TRYAGAIN redis error.
	retryConn, err := redisc.RetryConn(conn, 3, 100*time.Millisecond)
	if err != nil {
		log.Fatalf("RetryConn failed: %v", err)
	}

	//// set
	//for i := 0; i < n; i++ {
	//	key := fmt.Sprintf("some-key%d", i)
	//	val := fmt.Sprintf("val%d", i)
	//	_, err := retryConn.Do("SET", key, val)
	//	if err != nil {
	//		log.Fatalf("SET failed: %v", err)
	//	}
	//}

	// get
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("some-key%d", i)
		val := fmt.Sprintf("val%d", i)
		r, err := redis.String(retryConn.Do("GET", key))
		if err != nil {
			log.Fatalf("SET failed: %v", err)
		}
		assert.Equal(t, r, val)
	}
}
