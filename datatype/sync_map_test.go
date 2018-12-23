// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"strconv"
	"sync"
	"testing"
)

// region different map implementations

// see https://blog.golang.org/go-maps-in-action
// And now we’ve arrived as to why the sync.Map was created. The Go team identified situations in the standard lib
// where performance wasn’t great. There were cases where items were fetched from data structures wrapped in a sync.RWMutex,
// under high read scenarios// while deployed on very high multi-core setups and performance suffered considerably.

type counterMap struct {
	data map[string]int
}

func (m *counterMap) read(k string) int {
	return m.data[k]
}

func (m *counterMap) write(k string, v int) {
	m.data[k] = v
}

func (m *counterMap) inc(k string) {
	// info: need lock, check TestSyncMap
	m.data[k] = m.data[k] + 1
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type counterSafeMap struct {
	sync.RWMutex
	data map[string]int
}

func (m *counterSafeMap) read(k string) int {
	m.RLock()
	defer m.RUnlock()
	return m.data[k]
}

func (m *counterSafeMap) write(k string, v int) {
	m.Lock()
	defer m.Unlock()
	m.data[k] = v
}

func (m *counterSafeMap) inc(k string) {
	m.Lock()
	defer m.Unlock()
	m.data[k] = m.data[k] + 1
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type counterSyncMap struct {
	data sync.Map
}

func (m *counterSyncMap) read(k string) (int, bool) {
	if v, ok := m.data.Load(k); ok {
		return v.(int), true
	}

	return -1, false
}

func (m *counterSyncMap) write(k string, v int) {
	m.data.Store(k, v)
}

func (m *counterSyncMap) inc(k string) {
	// info: need lock, check TestSyncMap
	if v, ok := m.data.Load(k); ok {
		m.data.Store(k, v.(int)+1)
		return
	}

	m.data.Store(k, 1)
}

// endregion

// run several time and you may get result that the expected and result are different
// or get fatal error: concurrent map read and map write
func Test_Map(t *testing.T) {
	m := counterMap{data: make(map[string]int)}
	t.Log(m.read("10"))
	m.write("10", 10)
	t.Log(m.read("10"))

	var swg sync.WaitGroup

	for i := 0; i < 100; i++ {
		swg.Add(1)
		go func() {
			m.write("10", 10)
			swg.Done()
		}()
	}

	swg.Wait()
}

func Test_SafeMap(t *testing.T) {
	m := counterSafeMap{data: make(map[string]int)}
	t.Log(m.read("10"))
	m.write("10", 10)
	t.Log(m.read("10"))

	var swg sync.WaitGroup

	for i := 0; i < 100; i++ {
		swg.Add(1)
		go func() {
			m.write("10", 10)
			swg.Done()
		}()
	}

	swg.Wait()
}

func Test_SyncMap(t *testing.T) {
	m := counterSyncMap{}
	t.Log(m.read("10"))
	m.write("10", 10)
	t.Log(m.read("10"))

	var swg sync.WaitGroup

	for i := 0; i < 100; i++ {
		swg.Add(1)
		go func() {
			m.write("10", 10)
			swg.Done()
		}()
	}

	swg.Wait()
}

// 100000000	        18.7 ns/op	       0 B/op	       0 allocs/op
func Benchmark_WriteMap(b *testing.B) {
	m := counterMap{data: make(map[string]int)}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.write("10", 10)
	}
}

// 10000000	       174 ns/op	      40 B/op	       3 allocs/op
func Benchmark_WriteSyncMap(b *testing.B) {
	m := counterSyncMap{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.write("10", 10)
	}
}

// 50000000	        20.5 ns/op	       0 B/op	       0 allocs/op
func Benchmark_ReadMap(b *testing.B) {
	m := counterMap{data: make(map[string]int)}
	for i := 0; i < 100; i++ {
		m.write(strconv.Itoa(i), i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.read("10")
	}
}

// 20000000	        82.0 ns/op	       0 B/op	       0 allocs/op
func Benchmark_ReadSafeMap(b *testing.B) {
	m := counterSafeMap{data: make(map[string]int)}
	for i := 0; i < 100; i++ {
		m.write(strconv.Itoa(i), i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.read("10")
	}
}

// 30000000	        50.8 ns/op	       0 B/op	       0 allocs/op
func Benchmark_ReadSyncMap(b *testing.B) {
	m := counterSyncMap{}
	for i := 0; i < 100; i++ {
		m.write(strconv.Itoa(i), i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.read("10")
	}
}
