// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package main

import (
	"fmt"
	"github.com/tylertreat/BoomFilters"
)

func main() {
	cf := boom.NewCuckooFilter(1000, 0.01)

	cf.Add([]byte(`a`))
	if cf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if contains, _ := cf.TestAndAdd([]byte(`b`)); !contains {
		fmt.Println("doesn't contain b")
	}

	if cf.TestAndRemove([]byte(`b`)) {
		fmt.Println("removed b")
	}

	// Restore to initial state.
	cf.Reset()
}