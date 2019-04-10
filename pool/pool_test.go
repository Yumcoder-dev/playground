// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package pool

import (
	"fmt"
	"sync"
	"testing"
)

type A struct {
	name string
	age  int
}

func (a *A) Reset() {
	a.name = ""
	a.age = 0
}

func Test_pool(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return &A{}
		},
	}

	a := p.Get().(*A)
	a.name = "a"
	a.age = 10

	t.Log(a)

	p.Put(a)

	b := p.Get().(*A)
	t.Log(b)
	p.Put(b)

	b = p.Get().(*A)
	b.Reset()
	t.Log(b)
}

func f2(res []*A)  {
	for _, r := range res{
		if r.age > 0{
			fmt.Println(r.name)
		}
	}
}

func f1() []*A {
	res := make([]*A, 0)
	for i:= 0; i<100;i++ {
		res = append(res, &A{name:fmt.Sprintf("yumd%d", i), age:i})
	}
	return res
}

func Benchmark_pool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f2(f1())
	}
}

func f12(p *sync.Pool) []*A {
	res := make([]*A, 0)
	for i:= 0; i<100;i++ {
		obj := p.Get().(*A)
		obj.Reset()
		obj.name = fmt.Sprintf("yumd%d", i)
		obj.age = i
		res = append(res, obj)
	}
	return res
}

func Benchmark_pool2(b *testing.B) {
	p := &sync.Pool{
		New: func() interface{} {
			return &A{}
		},
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res := f12(p)
		f2(res)
		for _, r := range res{
			p.Put(r)
		}
	}
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func fStatic() []int32 {
	res := make([]int32, 2000)
	for i:=int32(0);i<2000;i++{
		res = append(res, i)
	}

	return res
}

func fDyanamic() []*int32 {
	res := make([]*int32, 2000)
	for i:=int32(0);i<2000;i++{
		res = append(res, &i)
	}

	return res
}

func Benchmark_fstatic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fStatic()
		//fDyanamic()
	}
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func fParameter(a, b, c int32, d, e, f int64, a1, a2, a3 string)  {
	fmt.Print(a, b, c)
	fmt.Print(d, e, f)
	fmt.Print(a1, a2, a3)
}

type p struct{
	a, b, c int32
	d, e, f int64
	a1, a2, a3 string
}
func fParameter2(i p)  {
	fmt.Print(i.a, i.b, i.c)
	fmt.Print(i.d, i.e, i.f)
	fmt.Print(i.a1, i.a2, i.a3)
}

func Benchmark_fParameter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fParameter(int32(1*i), int32(2*i), int32(3*i), 100, 200, 300, "a1","a2","a3")
	}
}

func Benchmark_fParameter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fParameter2(p{int32(1*i), int32(2*i), int32(3*i), 100, 200, 300, "a1","a2","a3"})
	}
}