// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package op

import (
	"math"
	"testing"
)

func Test_UintShift(t *testing.T) {
	u := uint64(2)
	i1 := int64(u | 1<<60)
	i2 := int64(u | 2<<60)
	t.Logf("%064b", i1)
	t.Logf("%064b", i2)

	t.Logf("%d", i2)
}

func Test_Shift(t *testing.T) {
	b := 12
	for (b & 1) == 0 {
		b = b >> 1
		t.Logf("%x", b)
	}

	t.Logf("%b", 1<<2)
	t.Logf("%b", 2<<2)

	t.Logf("%b", 1<<56)
	t.Logf("%b", 2<<56)

	t.Logf("%d", 1<<32)
	t.Logf("%f", math.Pow(2, 32))

}

func Test_Unset(t *testing.T) {
	t.Logf("%08b %d", 12&-4, (12&-4)%4)
	t.Logf("%08b %d", 13&-4, (13&-4)%4)
	t.Logf("%08b %d", 3&-4, (3&-4)%4)
	t.Logf("%08b %d", 5&-4, (5&-4)%4)

	t.Log((4 + 8 + 1) % 4)

}

func ifString(a string) bool {
	if a == "server" {
		return true
	}
	return false
}

func ifMath(a int32) bool {
	return a-10 == 0
}

func Benchmark_ifString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifString("yumcoder")
	}
}

func Benchmark_math(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifMath(100)
	}
}
