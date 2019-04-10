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
	sbf := boom.NewDefaultScalableBloomFilter(0.01)

	sbf.Add([]byte(`a`))
	if sbf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !sbf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if sbf.Test([]byte(`b`)) {
		fmt.Println("now it contains b!")
	}

	// Restore to initial state.
	sbf.Reset()
}