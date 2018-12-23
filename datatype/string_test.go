// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"testing"
)

// go test -bench=BenchmarkString* -benchmem

// region call_by_ref vs call_by_value
func strCallByValue(s string) string {
	s = "strCallByValue"
	return s
}

func strCallByRef(s *string) string {
	*s = "strCallByRef"
	return *s
}

func Test_StringCallByValue(t *testing.T) {
	var testTable = []struct {
		input    string
		expected string
	}{
		{"string", "string"},
		{"12345", "12345"},
		{"string2", "string2"},
	}
	for _, v := range testTable {
		result := strCallByValue(v.input)
		t.Log(result)
		if v.input != v.expected {
			t.Error("result: ", v.input, "expected: ", v.expected)
		}
	}
}

func Test_StringCallByRef(t *testing.T) {
	var testTable = []struct {
		input    string
		expected string
	}{
		{"string", "strCallByRef"},
		{"12345", "strCallByRef"},
		{"string2", "strCallByRef"},
	}
	for _, v := range testTable {
		result := strCallByRef(&v.input)
		t.Log(result)
		if v.input != v.expected {
			t.Error("result: ", v.input, "expected: ", v.expected)
		}
	}
}

var result string

func Benchmark_StringCallByValue(b *testing.B) {
	b.ReportAllocs()
	var r string

	for i := 0; i < b.N; i++ {
		s := "string"
		r = strCallByValue(s)
	}
	result = r
}

func Benchmark_StringCallByRef(b *testing.B) {
	b.ReportAllocs()
	var r string
	for i := 0; i < b.N; i++ {
		s := "string"
		r = strCallByRef(&s)
	}
	result = r
}

// endregion

func Test_BackQuote(t *testing.T) {
	ml := `
	a
	b
	`
	t.Log(`a\t\nb`)
	t.Log(ml)
}

func Test_DoubleQuote(t *testing.T) {
	t.Log("a\t\nb")
}

func Test_str(t *testing.T) {
	a := "omid\\gh"
	b := `omid\`
	t.Log(a)
	t.Log(b)
}
