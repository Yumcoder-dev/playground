// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package time

import (
	"testing"
	"time"
)

func Test_close_session(t *testing.T) {
	t1 := time.Now().Unix()
	t2 := t1 + 1
	t.Log(t1)
	t.Log(t2)
	time.Sleep(time.Second)
	t.Log(time.Now().Unix())
}

func Test_unix_to_time(t *testing.T) {
	utime := time.Unix(1556950953, 0)
	t.Log(utime.Format("02/01/2006, 15:04:05"))
}