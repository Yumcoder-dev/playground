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

// Top-K uses a Count-Min Sketch and min-heap to track the top-k most frequent elements in a stream
func main() {
	topk := boom.NewTopK(0.001, 0.99, 5)

	topk.Add([]byte(`bob`)).Add([]byte(`bob`)).Add([]byte(`bob`))
	topk.Add([]byte(`tyler`)).Add([]byte(`tyler`)).Add([]byte(`tyler`)).Add([]byte(`tyler`))
	topk.Add([]byte(`fred`))
	topk.Add([]byte(`alice`)).Add([]byte(`alice`)).Add([]byte(`alice`)).Add([]byte(`alice`))
	topk.Add([]byte(`james`))
	topk.Add([]byte(`fred`))
	topk.Add([]byte(`sara`)).Add([]byte(`sara`))
	topk.Add([]byte(`bill`))

	for i, element := range topk.Elements() {
		fmt.Println(i, string(element.Data), element.Freq)
	}

	// Restore to initial state.
	topk.Reset()
}
