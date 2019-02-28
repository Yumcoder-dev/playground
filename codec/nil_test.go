// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package codec

import "testing"

func Test_Nil_Arr(t *testing.T) {
	var m []int32
	m = nil
	t.Log(len(m))
	for i:= range m{
		t.Log(i)
	}
}
