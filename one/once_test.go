// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package one

import (
	"sync"
	"testing"
)

type onceFunCall int

func (o *onceFunCall) Increment() {
	*o++
}

func Test_once(t *testing.T) {
	var incFn onceFunCall
	once := &sync.Once{}
	once.Do(incFn.Increment)
	once.Do(incFn.Increment)
	once.Do(incFn.Increment)

	t.Log(incFn)
}
