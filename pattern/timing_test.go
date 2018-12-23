// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package pattern

import (
	"fmt"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	fmt.Printf("%s lasted %s", name, elapsed)
}

func Example_BigIntFactorial() {
	// Arguments to a defer statement is immediately evaluated and stored.
	// The deferred function receives the pre-evaluated values when its invoked.
	defer Duration(time.Now(), "IntFactorial")

	time.Sleep(5 * time.Second)

	// Output:
	// IntFactorial lasted 5.000114622s
}
