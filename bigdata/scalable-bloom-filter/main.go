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

// Stable Bloom Filters are useful for cases where the size of the data set isn't known a priori and memory is bounded.
// For example, an SBF can be used to deduplicate events from an unbounded event stream with a specified upper bound on
// false positives and minimal false negatives.
func main() {
	sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)
	fmt.Println("stable point", sbf.StablePoint())

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
