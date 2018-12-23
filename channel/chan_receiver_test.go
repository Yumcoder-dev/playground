// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package channel

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func onServer(i int) { println("S:", i) }
func onUser(i int)   { println("U:", i) }

func Test_Receiver(t *testing.T) {
	fromServer, fromUser := make(chan int), make(chan int)
	var serverData, userInput int
	var ok bool

	go func() {
		fromServer <- 1
		fromUser <- 1
		close(fromServer)
		runtime.Gosched()
		fromUser <- 2
		close(fromUser)
	}()

	isRunning := true
	for isRunning {
		select {
		case serverData, ok = <-fromServer:
			if ok {
				onServer(serverData)
			} else {
				isRunning = false
			}

		case userInput, ok = <-fromUser:
			if ok {
				onUser(userInput)
			} else {
				isRunning = false
			}
		}
	}
	println("end")
}

func Test_close_channel(t *testing.T) {
	fromServer, fromUser := make(chan int), make(chan int)
	w := &sync.WaitGroup{}
	w.Add(1)

	go func(wait *sync.WaitGroup) {
		defer func() {
			fmt.Println("end-go-routine")
			wait.Done()
		}()

		for {
			select {
			case <-fromServer:
				onServer(1)

			case <-fromUser:
				onUser(2)
				return
			}
		}
	}(w)

	fromServer <- 1
	close(fromUser)
	fmt.Println("end-main")
	w.Wait()
}

func Test_session_non_blocking_model(t *testing.T) {
	reqChan := make(chan int, 2)
	resChan := make(chan int, 2)

	worker := func(req int) int {
		//time.Sleep(time.Second)
		return req
	}
	go func() {
		for {
			select {
			case req := <-reqChan:
				fmt.Println("req--->", req)
				reply := worker(req)
				// resChan <- reply  --> all goroutines are asleep - deadlock!
				go func(reply int) {
					select {
					case resChan <- reply:
					}
				}(reply)
			case res := <-resChan:
				fmt.Println("res<-", res)
			case <-time.After(time.Second):
				fmt.Println("timer...")
			}
		}
	}()

	// sender
	go func() {
		fmt.Println("start sending...")
		for i := 0; i < 20; i++ {
			reqChan <- i
		}
		fmt.Println("sender quited...")
	}()

	select {}
}
