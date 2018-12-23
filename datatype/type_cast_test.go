// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"fmt"
	"testing"
)

type sss struct {
}

func f1(f interface{}) {
	if name, ok := f.(fmt.Stringer); !ok {
		fmt.Println(ok)
	} else {

		fmt.Println(name)
	}

}

func Test_C1(t *testing.T) {
	s := &sss{}
	f1(s)
}
