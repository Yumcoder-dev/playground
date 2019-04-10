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
	ibf := boom.NewInverseBloomFilter(10000)

	ibf.Add([]byte(`a`))
	if ibf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !ibf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if ibf.Test([]byte(`b`)) {
		fmt.Println("now it contains b!")
	}
}