// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package bench

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

// https://medium.com/@cep21/using-go-1-10-new-trace-features-to-debug-an-integration-test-1dc39e4e812d
// go test -cpuprofile cpu.prof
// go test -run=TestServer -cpuprofile cpu.prof
// tool pprof -http=localhost:12345 cpu.prof
// go test -run=TestServer -blockprofile=block.out
// go tool pprof -http=localhost:12345  block.out
// go test -run=TestServer -trace trace.out
// go tool trace trace.out
// http://127.0.0.1:42965/
// run on chrome and if trace_view is not show install go from source (misc folder must be in $GOROOT)
// todo: more read about practical trace
func takeCPU(start time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	j := 3
	for time.Since(start) < time.Second {
		for i := 1; i < 1000000; i++ {
			j *= i
		}
	}
	fmt.Println("j: ", j)
}

func takeTimeOnly(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
}

func takeIO(start time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	errCount := 0
	for time.Since(start) < time.Second*4 {
		_, err := http.Get("https://www.google.com")
		if err != nil {
			errCount++
		}

	}
	fmt.Println("io: ", errCount)
}

func Test_Server(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	start := time.Now()
	go takeCPU(start, &wg)
	go takeTimeOnly(&wg)
	go takeIO(start, &wg)
	wg.Wait()
}
