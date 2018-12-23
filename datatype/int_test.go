// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"math"
	"testing"
)

func Test_buf_decode(t *testing.T) {
	var a int32
	var b uint32

	a = 12
	b = uint32(a)
	t.Log(a)
	t.Log(b)

	a = -12
	b = uint32(a)
	t.Log(a)
	t.Log(b)
	t.Log((math.MaxInt32*2 + 1) - (12 - 1)) // [-127:127] --> [0:255] --> 127*2+1

	a = int32(b)
	t.Log(a)
}
