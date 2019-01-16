// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package recursion

import "testing"

type Result struct {
	ids []int
	errCode int32
	fType int32
}

func f1(a int, result *Result)  {
	switch a {
	case 1:
		for i:=0; i<10; i++{
			if result.errCode != 0{
				return
			}
			f1(a+2, result)
		}
	case 2:
		for i:=0; i<3; i++{
			f1(a+3, result)
		}
	case 3:
		result.fType = 100
		result.ids = append(result.ids, 3)
	case 4:
		result.fType = 200
		result.ids = append(result.ids, 4)
	case 5:
		result.fType = 200
		result.ids = append(result.ids, 5)
	case 6:
		result.fType = 100
		result.ids = append(result.ids, 6)
	case 7:
		result.fType = 400
		result.ids = append(result.ids, 7)
	case 8:
		result.fType = 500
		result.ids = append(result.ids, 8)
	case 9:
		result.fType = 500
		result.ids = append(result.ids, 9)
	default:
		result.fType = 100
		result.ids = append(result.ids, 10)
	}
}

func f2(a int ) (result *Result) {
	result = &Result{ids:make([]int, 0)}
	switch a {
	case 1:
		for i:=0; i<10; i++{
			if result.errCode != 0{
				return
			}
			r := f2(a+2)
			result.ids = append(result.ids, r.ids...)
			result.fType = r.fType
		}
	case 2:
		for i:=0; i<3; i++{
			r := f2(a+3)
			result.ids = append(result.ids, r.ids...)
			result.fType = r.fType
		}
	case 3:
		result.fType = 100
		result.ids = append(result.ids, 3)
	case 4:
		result.fType = 200
		result.ids = append(result.ids, 4)
	case 5:
		result.fType = 200
		result.ids = append(result.ids, 5)
	case 6:
		result.fType = 100
		result.ids = append(result.ids, 6)
	case 7:
		result.fType = 400
		result.ids = append(result.ids, 7)
	case 8:
		result.fType = 500
		result.ids = append(result.ids, 8)
	case 9:
		result.fType = 500
		result.ids = append(result.ids, 9)
	default:
		result.fType = 100
		result.ids = append(result.ids, 10)
	}

	return result
}

func f3(a int ) Result {
	result := Result{ids:make([]int, 0)}
	switch a {
	case 1:
		for i:=0; i<10; i++{
			if result.errCode != 0{
				return result
			}
			r := f3(a+2)
			result.ids = append(result.ids, r.ids...)
			result.fType = r.fType
		}
	case 2:
		for i:=0; i<3; i++{
			r := f3(a+3)
			result.ids = append(result.ids, r.ids...)
			result.fType = r.fType
		}
	case 3:
		result.fType = 100
		result.ids = append(result.ids, 3)
	case 4:
		result.fType = 200
		result.ids = append(result.ids, 4)
	case 5:
		result.fType = 200
		result.ids = append(result.ids, 5)
	case 6:
		result.fType = 100
		result.ids = append(result.ids, 6)
	case 7:
		result.fType = 400
		result.ids = append(result.ids, 7)
	case 8:
		result.fType = 500
		result.ids = append(result.ids, 8)
	case 9:
		result.fType = 500
		result.ids = append(result.ids, 9)
	default:
		result.fType = 100
		result.ids = append(result.ids, 10)
	}

	return result
}

func Benchmark_f1(b *testing.B) {
	result := &Result{ids:make([]int, 0)}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f1(1, result)
	}
	_ = result
}

func Benchmark_f2(b *testing.B) {
	var result *Result
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = f2(1)
	}
	_ = result
}

func Benchmark_f3(b *testing.B) {
	var result Result
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = f3(1)
	}
	_ = result
}