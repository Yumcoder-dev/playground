// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//

// example for register/deregister a server in etcd/consul/...
package pattern

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type done struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func newDone() *done {
	ctx, cancel := context.WithCancel(context.Background())
	d := &done{ctx: ctx, cancel: cancel}
	return d
}

func (r *done) register() {
	ticker := time.NewTicker(time.Second)
L:
	for {
		select {
		case <-r.ctx.Done():
			fmt.Println("done register...")
			break L
		case <-ticker.C:
			// etcd write
			fmt.Println("register...")
		}
	}
	// clean/release resource
	fmt.Println("end register...")
	defer r.wg.Done()
}

func (r *done) deregister(wg *sync.WaitGroup) {
	// context's Done channel is closed when cancel function is called
	r.cancel()
	r.wg = wg
}

func Test_Context(t *testing.T) {
	d := newDone()
	go d.register()
	time.Sleep(2. * time.Second)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	d.deregister(wg)
	wg.Wait()
	fmt.Println("end test...")
}
