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

// Count-Min Sketches are useful for counting the frequency of events in massive data sets or unbounded streams online.
// In these situations, storing the entire data set or allocating counters for every event in memory is impractical.
// It may be possible for offline processing, but real-time processing requires fast, space-efficient solutions like the CMS
func main() {
	cms := boom.NewCountMinSketch(0.001, 0.99)

	cms.Add([]byte(`alice`)).Add([]byte(`bob`)).Add([]byte(`bob`)).Add([]byte(`frank`))
	fmt.Println("frequency of alice", cms.Count([]byte(`alice`)))
	fmt.Println("frequency of bob", cms.Count([]byte(`bob`)))
	fmt.Println("frequency of frank", cms.Count([]byte(`frank`)))


	// Serialization example
	buf := new(bytes.Buffer)
	n, err := cms.WriteDataTo(buf)
	if err != nil {
		fmt.Println(err, n)
	}

	// Restore to initial state.
	cms.Reset()

	newCMS := boom.NewCountMinSketch(0.001, 0.99)
	n, err = newCMS.ReadDataFrom(buf)
	if err != nil {
		fmt.Println(err, n)
	}

	fmt.Println("frequency of frank", newCMS.Count([]byte(`frank`)))


}
