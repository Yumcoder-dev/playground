// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package proc

import (
	"crypto/rand"
	"testing"
)

func fval(a int, b int) int{
	t:= 0
	for i:=0; i<1000;i++{
		t += a+b
	}
	return t
}

func fref(a *int, b *int) int{
	t:= 0
	for i:=0; i<1000;i++{
		t += *a + *b
	}
	return t
}

func Test_fval(t *testing.T) {
	a := 10
	b := 20
	t.Log(fval(a, b))
}

func Test_fref(t *testing.T) {
	a := 10
	b := 20
	t.Log(fref(&a, &b))
}

func Benchmark_fval(b *testing.B) {
	p := 10
	q := 20
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fval(p, q)
	}
}

func Benchmark_fref(b *testing.B) {
	p := 10
	q := 20
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fref(&p, &q)
	}
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Data struct {
	a []byte
	b int
}

func f1(a []byte, b int){
	for i:=0; i<1000;i++{
		a[i] = a[i] + 10
	}
}

func f2(d *Data){
	for i:=0; i<1000;i++{
		d.a[i] = d.a[i] + 10
	}
}

func f3(d Data){
	for i:=0; i<1000;i++{
		d.a[i] = d.a[i] + 10
	}
}

func Benchmark_f1(b *testing.B) {
	a := make([]byte, 100000)
	rand.Read(a)

	for i := 0; i < b.N; i++ {
		f1(a, 100)
	}
}

func Benchmark_f2(b *testing.B) {
	a := make([]byte, 100000)
	rand.Read(a)
	d := &Data{a: a, b : 100}
	for i := 0; i < b.N; i++ {
		f2(d)
	}
}

func Benchmark_f3(b *testing.B) {
	a := make([]byte, 100000)
	rand.Read(a)
	d := Data{a: a, b : 100}
	for i := 0; i < b.N; i++ {
		f3(d)
	}
}