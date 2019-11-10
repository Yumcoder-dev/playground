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
	TLConstructor := -1721631396
	fmt.Printf("%x", uint32(TLConstructor))
}

func Test_ToHexadecimal(t *testing.T) {
	i := uint64(12240908862933197005)
	fmt.Printf("%x\n", i)
}
func setF ( s []int) []int{
	s[0] = len(s)
	return s
}
func Test_t(t *testing.T) {
	a := make([]int, 4, 5)
	b := setF(a)
	a[1]= 1
	b = append(b, 5)
	b[3]=3
	fmt.Println(a)
	fmt.Println(b)

}