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
	bf := boom.NewDefaultCountingBloomFilter(1000, 0.01)

	bf.Add([]byte(`a`))
	if bf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !bf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if bf.TestAndRemove([]byte(`b`)) {
		fmt.Println("removed b")
	}

	// Restore to initial state.
	bf.Reset()
}