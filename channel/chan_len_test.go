// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package channel

import "testing"

func Test_ChanLen(t *testing.T) {
	ch := make(chan int, 100) // size is mandatory
	ch <- 1
	ch <- 2
	for len(ch) > 0 {
		t.Log(<-ch) // 1, 2
	}

	t.Log(len(ch)) // 0
	t.Log(cap(ch)) // 100

	chs := make(chan struct{}, 1)
	chs <- struct{}{}
	t.Log(len(chs)) // 1
	t.Log(cap(chs)) // 1
}
