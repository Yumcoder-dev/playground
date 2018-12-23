// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package regex

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_TLC_ExtractLayer(t *testing.T) {
	//match := regexp.MustCompile("// LAYER( \\((\\w*)\\) | )(\\d+)").FindStringSubmatch("// LAYER (Ios) 11")
	//fmt.Println(`layer` + match[2] + match[3])

	match2 := regexp.MustCompile("// LAYER( \\((\\w*)\\) | )(\\d+)").FindStringSubmatch("// LAYER 11\n")
	fmt.Println(`layer` + match2[2] + match2[3])
	// Output:
	// true
	// true
	// false
	// false
}
