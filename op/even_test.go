// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package op

import "testing"

func Benchmark_even(b *testing.B) {
	tests := []struct {
		Name   string
		Action func(i int) bool
	}{
		{
			Name: "even",
			Action: func(i int) bool {
				return i%2 == 0
			},
		},
		{
			Name: "dr_even",
			Action: func(i int) bool {
				for {
					if i == 0 {
						return true
					} else if i == 1 {
						return false
					}
					i *= i
				}
			},
		},
	}
	for _, tc := range tests {
		b.Run(tc.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.Action(i)
			}
		})
	}
}
