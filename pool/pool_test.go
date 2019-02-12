// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package pool

import (
	"sync"
	"testing"
)

type A struct {
	name string
	age  int
}

func (a *A) Reset() {
	a.name = ""
	a.age = 0
}

func Test_poll(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return &A{}
		},
	}

	a := p.Get().(*A)
	a.name = "a"
	a.age = 10

	t.Log(a)

	p.Put(a)

	b := p.Get().(*A)
	t.Log(b)
	p.Put(b)

	b = p.Get().(*A)
	b.Reset()
	t.Log(b)
}
