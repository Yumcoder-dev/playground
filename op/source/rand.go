// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package source

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("init")
}

func Print()  {
	fmt.Println("Print")
}
