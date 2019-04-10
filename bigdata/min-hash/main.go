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

// MinHash is a probabilistic algorithm which can be used to cluster or compare documents by splitting the corpus into
// a bag of words. MinHash returns the approximated similarity ratio of the two bags. The similarity is less accurate
// for very small bags of words.
func main() {
	bag1 := []string{"bill", "alice", "frank", "bob", "sara", "tyler", "james"}
	bag2 := []string{"bill", "alice", "frank", "bob", "sara"}

	fmt.Println("similarity", boom.MinHash(bag1, bag2))
}
