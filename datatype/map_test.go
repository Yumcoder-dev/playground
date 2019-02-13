// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"google.golang.org/grpc/resolver"
	"testing"
)

// region map pass by ref
func mapPassByRef(s map[int]int, d int) {
	s[d] = d
}

func Test_ArgMap(t *testing.T) {
	s := make(map[int]int)
	s[10] = 10
	mapPassByRef(s, 20)
	for k, v := range s {
		t.Log("key: ", k, " value: ", v)
	}
}
func Benchmark_MapPassByRef(b *testing.B) {
	s := make(map[int]int)
	s[10] = 10
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mapPassByRef(s, 20)
	}
}

// endregion

// region unique structure
func appendUniqueSlice(a []int, x int) []int {
	for _, y := range a {
		if x == y {
			return a
		}
	}
	return append(a, x)
}

func appendUniqueMap(a map[int]int, x int) map[int]int {
	if _, ok := a[x]; ok {
		return a
	}

	a[x] = x
	return a
}

func getFromSlice(a []int, x int) int {
	for _, y := range a {
		if x == y {
			return y
		}
	}
	return -1
}

func getFromMap(a map[int]int, x int) int {
	return a[x]
}

func Test_AppendUniqueSlice(t *testing.T) {
	testTable := []struct {
		value int
	}{
		{10},
		{20},
		{30},
		{10},
	}
	s := make([]int, 0)
	for _, r := range testTable {
		s = appendUniqueSlice(s, r.value)
	}

	for k, v := range s {
		t.Log("key: ", k, " value: ", v)
	}
}

func Test_AppendUniqueMap(t *testing.T) {
	testTable := []struct {
		value int
	}{
		{10},
		{20},
		{30},
		{10},
	}
	s := make(map[int]int)
	s[10] = 10
	for _, r := range testTable {
		appendUniqueMap(s, r.value)
	}

	for k, v := range s {
		t.Log("key: ", k, " value: ", v)
	}
}

// 300000000	         5.12 ns/op	       0 B/op	       0 allocs/op
// 100000000	        10.8 ns/op	       0 B/op	       0 allocs/op
func Benchmark_AppendUniqueMap(b *testing.B) {
	s := make(map[int]int)
	for i := 0; i < 10000; i++ {
		appendUniqueMap(s, i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		appendUniqueMap(s, 10)
	}
}

// 300000000	         4.26 ns/op	       0 B/op	       0 allocs/op
// 200000000	         9.02 ns/op	       0 B/op	       0 allocs/op
func Benchmark_AppendUniqueSlice(b *testing.B) {
	s := make([]int, 0)
	for i := 0; i < 10000; i++ {
		s = appendUniqueSlice(s, i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = appendUniqueSlice(s, 10)
	}
}

// 100000000	        17.3 ns/op	       0 B/op	       0 allocs/op
func Benchmark_GetFromMap(b *testing.B) {
	s := make(map[int]int)
	for i := 0; i < 10000; i++ {
		appendUniqueMap(s, i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		getFromMap(s, -99)
	}
}

// 500000	      3763 ns/op	       0 B/op	       0 allocs/op
func Benchmark_GetFromSlice(b *testing.B) {
	s := make([]int, 0)
	for i := 0; i < 10000; i++ {
		s = appendUniqueSlice(s, i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		getFromSlice(s, -99)
	}
}

// endregion

func Test_MapStruct(t *testing.T) {
	state := make(map[string]struct{})
	state["yumcoder"] = struct{}{}
	_, ok := state["yumcoder2"]
	t.Log(ok)
}

func Test_Map_Mack(t *testing.T) {
	v := make(map[string]map[int]bool)
	_, ok := v["a"][1]
	if !ok {
		v["a"] = make(map[int]bool)
	}
	v["a"][1] = true
	t.Log(v["a"][1])
}

func Test_MapOfStruct(t *testing.T) {
	type Dummy struct {
		a int
	}

	//addrs  := []resolver.Address {
	//	{Addr:"a", Metadata:Dummy{150}}, // panic: runtime error: hash of unhashable type map[string]interface {}
	//	{Addr:"b"},
	//}

	addrs := []resolver.Address{
		{Addr: "a", Metadata: &Dummy{150}},
		{Addr: "b"},
	}

	addrsSet := make(map[resolver.Address]struct{})
	for _, a := range addrs {
		addrsSet[a] = struct{}{}
	}

	t.Log(addrsSet)
}

func Test_add_to_map(t *testing.T) {
	sessionByType := make(map[int8][]int, 0)
	data := []struct{
		t int8
		val int
	}{
		{1, 1},{1,2}, {2,3},{3, 3}, {3, 4},
	}

	for index, userSession := range data {
		if v, ok := sessionByType[userSession.t]; ok {
			v = append(v, index)
			sessionByType[userSession.t] = v
		}else{
			sessionByType[userSession.t] = []int{index}
		}
	}

	t.Log(sessionByType)
}
