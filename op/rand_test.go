// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package op

import (
	"math/rand"
	"testing"
	"yumcoder.com/playground/op/source"
)

func Test_read(t *testing.T) {
	source.Print()
	handshake := make([]byte, 16)
	rand.Read(handshake)
	t.Log(handshake)
}
func Test_rand(t *testing.T) {
	source.Print()

	t.Log(rand.Int63())
	t.Log(rand.Int63())
}
