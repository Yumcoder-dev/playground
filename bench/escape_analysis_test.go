// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package bench

import (
	"bytes"
	"io"
	"testing"
)

// region simple example
/*
sum adds the numbers between 1 and 100 and returns the result. This is a rather unusual way to do this, but it
illustrates how Escape Analysis works. Because the numbers slice is only referenced inside sum, the compiler will
arrange to store the 100 integers for that slice on the stack, rather than the heap. There is no need to garbage
collect numbers, it is automatically freed when sum returns.
*/
func sum(size int) int {
	numbers := make([]int, size) // never scape sum()
	for i := range numbers {
		numbers[i] = i + 1
	}
	var sum int
	for i := range numbers {
		sum += i
	}

	return sum
}

// 2000000000	         0.00 ns/op	       0 B/op	        0 allocs/op --> never scape and the size is limited
// 2000000000	         0.00 ns/op	       0 B/op	        0 allocs/op --> never scape and the size is limited
// 1	                 2666452818 ns/op  4194304096 B/op	2 allocs/op --> too large for stack
func Benchmark_Sum(b *testing.B) {
	benchTable := []struct {
		name string
		size int
	}{
		{"size100", 100},
		{"size100k", 100 * 1024},
		{"size100M", 500 * 1024 * 1024},
	}
	for _, v := range benchTable {
		b.Run(v.name, func(b *testing.B) {
			b.ReportAllocs()
			sum(v.size)
		})
	}
}

// endregion

// region inline and interface example
type interfaceScapeToHeap interface {
	chaneName()
}

type user struct {
	name  string
	email string
}

func (u *user) chaneName() {
	u.name = "name is changed by calling interface"
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "createUserV1",
		email: "noinlineV1@user.com",
	}

	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "createUserV2",
		email: "noinlineV2@user.com",
	}

	return &u
}

func createUserV3() *user { // inline func
	u := user{
		name:  "createUserV3",
		email: "inlineV3@user.com",
	}

	return &u
}

func consumeUser(u *user) {
	if len(u.name) > 0 {
		u.name = "new name"
	}
}

func callInterface(i interfaceScapeToHeap) {
	i.chaneName()
}

// 1000000000	         2.43 ns/op	       0 B/op	       0 allocs/op
func Benchmark_CreateUserV1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createUserV1()
	}
}

// 30000000	        48.7 ns/op	      32 B/op	       1 allocs/op
// note: if remove //go:noinline for createUserV2 then allocation to be 0 because in that
// situation inline func is used and user does not exposed to BenchmarkCreateUserV2()
func Benchmark_CreateUserV2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createUserV2() // 1 allocs/op --> user expose outside of createUserV2()
	}
}

// 1000000000	         2.67 ns/op	       0 B/op	       0 allocs/op
func Benchmark_CreateUserV3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		u2 := createUserV3() // inline --> and called in next method, so it can create on stack
		consumeUser(u2)
	}
}

// 30000000	        50.0 ns/op	      32 B/op	       1 allocs/op
func Benchmark_CreateUserV4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		u2 := createUserV3() // inline --> pass as interface
		callInterface(u2)
	}
}

//endregion

// region memory analysis

//also see // https://www.ardanlabs.com/blog/2017/06/language-mechanics-on-memory-profiling.html
//compiler reporting
//$ go build -gcflags "-m -m"

// region algorithm

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// algOne is one way to solve the problem.
func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {

		// Read in one byte from the input stream.
		if _, err := io.ReadFull(input, buf[end:]); err != nil {

			// Flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// If we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// Read a new initial number of bytes.
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// Write the gateway byte since it has been compared.
		output.WriteByte(buf[0])

		// Slice that gateway byte out.
		copy(buf, buf[1:])
	}
}

// algTwo is a second way to solve the problem.
// Provided by Tyler Bunnell https://twitter.com/TylerJBunnell
func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use the bytes Reader to provide a stream to process.
	input := bytes.NewReader(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Create an index variable to match bytes.
	idx := 0

	for {

		// Read a single byte from our input.
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		// Does this byte match the byte at this offset?
		if b == find[idx] {

			// It matches so increment the index position.
			idx++

			// If every byte has been matched, write
			// out the replacement.
			if idx == size {
				output.Write(repl)
				idx = 0
			}

			continue
		}

		// Did we have any sort of match on any given byte?
		if idx != 0 {

			// Write what we've matched up to this point.
			output.Write(find[:idx])

			// Unread the unmatched byte so it can be processed again.
			input.UnreadByte()

			// Reset the offset to start matching from the beginning.
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset.
		output.WriteByte(b)
		idx = 0
	}
}

// assembleInputStream combines all the input into a
// single stream for processing.
func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

// endregion

// profiling:
// go test -run none -bench AlgorithmOne -benchtime 3s -benchmem -memprofile mem.out
// BenchmarkAlgorithmOne-8          1000000              3024 ns/op             117 B/op          2 allocs/op
// --> 2 alloc --> want to improve performance
// go tool pprof -alloc_space memcpu.test mem.out
// (pprof) list algOne
// Total: 110.51MB
// ROUTINE ======================== yumcoder.com/benchmark ....
//       3MB   104.01MB (flat, cum)   100% of Total
//         .          .     35:
//         .          .     36:// algOne is one way to solve the problem.
//         .          .     37:func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
//         .          .     38:
//         .          .     39:   // Use a bytes Buffer to provide a stream to process.
//         .   101.01MB     40:   input := bytes.NewBuffer(data)
//         .          .     41:
//         .          .     42:   // The number of bytes we are looking for.
//         .          .     43:   size := len(find)
//         .          .     44:
//         .          .     45:   // Declare the buffers we need to process the stream.
//       3MB        3MB     46:   buf := make([]byte, size)
//         .          .     47:   end := size - 1
//         .          .     48:
//         .          .     49:   // Read in an initial number of bytes we need to get started.
//         .          .     50:   if n, err := io.ReadFull(input, buf[:end]); err != nil {
//         .          .     51:           output.Write(buf[:n])
//
//
// Based on this profile, we now know input and the making array of the buf slice is allocating to the heap.
// ------- line 40 and 46 -------
// We could assume it is allocating because the function call to bytes.NewBuffer is sharing the bytes.Buffer value it
// creates up the call stack. However, if the existence of a value in the flat column (the first column in the pprof output)
// tells me that the value is allocating because the algOne function is sharing it in a way to cause it to escape.
// anyway run $ go build -gcflags "-m -m"
// line 40 in benchmark.go input := bytes.NewBuffer(data) so we concentrate on it, result show bellow
//
//	./benchmark.go:40:26: inlining call to bytes.NewBuffer func([]byte) *bytes.Buffer { return &bytes.Buffer literal }
//	./benchmark.go:40:26: &bytes.Buffer literal escapes to heap    <-----------
//	./benchmark.go:40:26:   from ~R0 (assign-pair) at ./benchmark.go:40:26
//	./benchmark.go:40:26:   from input (assigned) at ./benchmark.go:40:8
//	./benchmark.go:40:26:   from input (interface-converted) at ./benchmark.go:50:26     <-----------
//	./benchmark.go:40:26:   from input (passed to call[argument escapes]) at ./benchmark.go:50:26
//
//	./benchmark.go:46:13: make([]byte, size) escapes to heap
//	./benchmark.go:46:13:   from make([]byte, size) (too large for stack) at ./benchmark.go:46:13
//
// The report says the creating array is “too large for stack”. This message is very misleading.
// It’s not that the backing array is too large, but that the compiler doesn’t know what the size is in compile time.
func Benchmark_AlgorithmOne(b *testing.B) {
	b.ReportAllocs()
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// 2000000	       739 ns/op	       0 B/op	       0 allocs/op
func Benchmark_AlgorithmTwo(b *testing.B) {
	b.ReportAllocs()
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}

// endregion
