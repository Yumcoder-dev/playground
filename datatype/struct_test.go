// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// region struct call by value
type person struct {
	firstName string
	lastName  string
}

func changeName(p person) {
	p.firstName = "bug"
}

func Test_StructCallByValue(t *testing.T) {
	person := person{
		firstName: "yumcoder",
		lastName:  "developer",
	}

	changeName(person)

	t.Log(person)
	assert.Equal(t, person.firstName, "yumcoder")
}

// endregion

// region update a slice of struct
type E struct {
	A, B, C, D int
}

func (e *E) update(a, b, c, d int) {
	e.A += a
	e.B += b
	e.C += c
	e.D += d
}

var SIZE = 1000

// 1000000	      1517 ns/op	       0 B/op	       0 allocs/op
func Benchmark_SliceOfStructUpdate(b *testing.B) {
	b.ReportAllocs()
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			e[i].update(1, 2, 3, 4)
		}
	}
	b.StopTimer()
}

// 1000000	      1490 ns/op	       0 B/op	       0 allocs/op
func Benchmark_SliceOfStructUpdateManual(b *testing.B) {
	b.ReportAllocs()
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			e[i].A += 1
			e[i].B += 2
			e[i].C += 3
			e[i].D += 4
		}
	}
	b.StopTimer()
}

// 1000000	      1485 ns/op	       0 B/op	       0 allocs/op
func Benchmark_SliceOfStructUpdateUnroll(b *testing.B) {
	b.ReportAllocs()
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i := range e {
			v := &e[i]
			v.A += 1
			v.B += 2
			v.C += 3
			v.D += 4
		}
	}
	b.StopTimer()
}

// 1000000	      1476 ns/op	       0 B/op	       0 allocs/op
func Benchmark_SliceOfStructUpdateRange(b *testing.B) {
	b.ReportAllocs()
	var e = make([]E, SIZE)
	for j := 0; j < b.N; j++ {
		for i, v := range e {
			v.A += 1
			v.B += 2
			v.C += 3
			v.D += 4
			e[i] = v
		}
	}
	b.StopTimer()
}

// endregion
