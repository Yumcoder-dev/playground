// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//

// test running statement oder in select
package channel

import "testing"

func Test_select(t *testing.T) {
	//ch1 := make(chan interface{}, 1)
	//go func() {
	//	select {
	//	case ch1 <-nil:
	//	}
	//
	//}()
	//
	//
	//<-ch1

	ch2 := make(chan interface{}, 10)
	ch3 := make(chan interface{}, 10)
	ch2 <- "m2"
	ch2 <- "m22"
	ch3 <- "m3"

	select {
	case m2 := <-ch2:
		t.Log(m2)
	}

	select {
	case m2 := <-ch2:
		t.Log(m2)
	case m2 := <-ch3:
		t.Log(m2)
	}

	// may m2 or m3
}

func sender() (chan int, chan bool) {
	worker := func(out chan int, exit chan bool) {
		for i := 1; i <= 10; i++ {
			out <- i
		}
		exit <- true
	}
	out := make(chan int, 10)
	exit := make(chan bool)
	go worker(out, exit)
	return out, exit
}

func Test_Order_SelectOnly(t *testing.T) {
	out, exit := sender()

L:
	for {
		select {
		case val := <-out:
			t.Logf("value: %d\n", val)
			continue
		case <-exit:
			t.Log("exiting")
			break L
		}
	}
	t.Log("did we get all 10? Most likely not")
}

func Test_Order_WithPriority(t *testing.T) {
	out, exit := sender()

L:
	for {

		select {
		case val := <-out:
			t.Logf("value: %d\n", val)
			continue
		default:
		}

		select {
		case val := <-out:
			t.Logf("value: %d\n", val)
			continue
		case <-exit:
			t.Log("exiting")
			break L
		}
	}
	t.Log("did we get all 10? I think so!")
}

func Test_Order_SeparateSelect(t *testing.T) {
	out, exit := sender()

L:
	for {
		select {
		case <-exit:
			t.Log("exiting")
			break L
		default:
		}

		select {
		case val := <-out:
			t.Logf("value: %d\n", val)
		default:
		}
	}
	t.Log("did we get all 10? I think so!")
}

func Test_Order_SeparateSelect_2(t *testing.T) {
	out, exit := sender()

L:
	for {

		select {
		case val := <-out:
			t.Logf("value: %d\n", val)
		default:
		}

		select {
		case <-exit:
			t.Log("exiting")
			break L
		default:
		}
	}
	t.Log("did we get all 10? I think so!")
}
