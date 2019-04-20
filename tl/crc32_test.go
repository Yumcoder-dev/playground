// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package tl

import (
	"fmt"
	"testing"
)

func Test_TLConstructorCrc32ToHexadecimal(t *testing.T) {
	TLConstructor := -1640190800
	fmt.Printf("%x", uint32(TLConstructor))
}

func Test_ToHexadecimal(t *testing.T) {
	i := uint64(12240908862933197005)
	fmt.Printf("%x\n", i)
}
