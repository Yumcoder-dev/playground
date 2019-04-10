// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package main

import (
	"bytes"
	"fmt"
	"github.com/tylertreat/BoomFilters"
)

// HyperLogLog is a probabilistic algorithm which approximates the number of distinct elements in a multiset
func main() {
	hll, err := boom.NewDefaultHyperLogLog(0.1)
	if err != nil {
		panic(err)
	}

	hll.Add([]byte(`alice`)).Add([]byte(`bob`)).Add([]byte(`bob`)).Add([]byte(`frank`))
	fmt.Println("count", hll.Count())

	// Serialization example
	buf := new(bytes.Buffer)
	_, err = hll.WriteDataTo(buf)
	if err != nil {
		fmt.Println(err)
	}

	// Restore to initial state.
	hll.Reset()

	newHll, err := boom.NewDefaultHyperLogLog(0.1)
	if err != nil {
		fmt.Println(err)
	}

	_, err = newHll.ReadDataFrom(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("count", newHll.Count())

}
