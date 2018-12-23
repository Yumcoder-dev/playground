// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package bench

import (
	"bytes"
	"html/template"
	"testing"
)

// see https://blog.golang.org/profiling-go-programs

// go test -bench=.
// go test -bench=. -benchmem
// result meaning
// ----------------------------------------------------------------
// BenchmarkParseTemplate-8          100000             24672 ns/op            9893 B/op         80 allocs/op
// BenchmarkExecuteTemplate-8       1000000              1954 ns/op             504 B/op         12 allocs/op
// ----------------------------------------------------------------
// 100000 is the number of iterations for i := 0; i < b.N; i++ {}
// XXX ns/op is approximate time it took for one iteration to complete
// allocs/op means how many distinct memory allocations occurred per op (single iteration).
// B/op is how many average bytes were allocated per op.
// meaning of BenchmarkXXX-8 means 8 core
// -- if go test -bench=. -benchmem -cpu=2 --> BenchmarkXXX-2
// -- if go test -bench=. -benchmem -cpu=4 --> BenchmarkXXX-4

// for test you must disable CPU frequency scaling
// meaning of CPU frequency scaling
// -- enables the operating system to scale the CPU frequency up or down in order to save power
// todo: disable CPU frequency scaling

// ------------------ benchmark notes ------------------
// Each benchmark is run for a minimum of 1 second by default.
// If the second has not elapsed when the Benchmark function returns,
// the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, …
// and the function run again.
// to define min time `$ go test -bench=. -benchtime=20s`

// ------------------ Traps for young players ------------------
//func BenchmarkWrong1(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		F1(n)
//	}
//}
//
//func BenchmarkWrong2(b *testing.B) {
//	F1(b.N)
//}
//
//func benchmarkCorrect(i int, b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		F1(i)
//	}
//}
//
//func BenchmarkC1(b *testing.B)  { benchmarkCorrect(1, b) }
//func BenchmarkC2(b *testing.B)  { benchmarkCorrect(2, b) }
//func BenchmarkC3(b *testing.B)  { benchmarkCorrect(3, b) }
//// ----- alternatively -----
//func BenchmarkC1(b *testing.B){
//	var benches = []struct {
//		name    string
//		n  int
//	}{
//		{"b1", 1},
//		{"b2", 2},
//		{"b3", 3},
//	}
//
//	for _, c := range benches {
//		b.Run(c.name, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				F1(c.N)
//			}
//		})
//	}
//}

// analyse the CPU profile
// 1. $ go test -bench=. -cpuprofile=cpu.out
// analyse the memory profile
// 1.1 $ go test -bench=. -memprofile=mem.out
// 2. $ go tool pprof bench.test cpu.out
// You can check the top20 consuming functions
// 3. (pprof) top20
// see the CPU consumption by cumulative time
// 4. (pprof) top --cum
// stack trace show result in form of a image
// 5.  (pprof) web
// 5.1 if graphviz is not install run the following `$ sudo apt install graphviz`
// -- Each box in the graph corresponds to a single function, and the boxes are sized according to the number of samples
// in which the function was running
// -- An edge from box X to box Y indicates that X calls Y;
// the number along the edge is the number of times that call appears in a sample

/*
Instrumentation is the process of adding code to your application to generate events to allow you to monitor application
health and performance. Instrumentation allows you to profile applications. Profiling enables you to identify how long
a particular method or operation takes to run and how efficient it is in terms of CPU and memory resource usage

(pprof) top10
Total: 2525 samples
     298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
     268  10.6%  22.4%     2124  84.1% main.FindLoops
     251   9.9%  32.4%      451  17.9% scanblock
     178   7.0%  39.4%      351  13.9% hash_insert
     131   5.2%  44.6%      158   6.3% sweepspan
     119   4.7%  49.3%      350  13.9% main.DFS
      96   3.8%  53.1%       98   3.9% flushptrbuf
      95   3.8%  56.9%       95   3.8% runtime.aeshash64
      95   3.8%  60.6%      101   4.0% runtime.settype_flush
      88   3.5%  64.1%      988  39.1% runtime.mallocgc

When CPU profiling is enabled, the Go program stops about 100 times per second and records a sample consisting of
the program counters on the currently executing goroutine's stack. The profile has 2525 samples, so it was running for
a bit over 25 seconds. In the `go tool pprof` output, there is a row for each function that appeared in a sample.

The first two columns show the number of samples in which the function was running (as opposed to waiting for a called
function to return), as a raw count and as a percentage of total samples. The runtime.mapaccess1_fast64 function was
running during 298 samples, or 11.8%. The top10 output is sorted by this sample count.

The third column shows the running total during the listing: the first three rows account for 32.4% of the samples.
cumulative of the second column: (11.8% + 10.6%) = 22.4%

The fourth and fifth columns show the number of samples in which the function appeared (either running or waiting for
a called function to return). The main.FindLoops function was running in 10.6% of the samples, but it was on the call
stack (it or functions it called were running) in 84.1% of the samples.

To sort by the fourth and fifth columns, use the -cum (for cumulative) flag:

(pprof) top5 -cum
Total: 2525 samples
       0   0.0%   0.0%     2144  84.9% gosched0
       0   0.0%   0.0%     2144  84.9% main.main
       0   0.0%   0.0%     2144  84.9% runtime.main
       0   0.0%   0.0%     2124  84.1% main.FindHavlakLoops
     268  10.6%  10.6%     2124  84.1% main.FindLoops

performance analysis for a function:
(pprof) list function_name
Total: 2525 samples
ROUTINE ====================== main.DFS in /home/rsc/g/benchgraffiti/havlak/havlak1.go
   119    697 Total samples (flat / cumulative)
     3      3  240: func DFS(currentNode *BasicBlock, nodes []*UnionFindNode, number map[*BasicBlock]int, last []int, current int) int {
     1      1  241:     nodes[current].Init(currentNode, current)
     1     37  242:     number[currentNode] = current
     .      .  243:
     1      1  244:     lastid := current
    89     89  245:     for _, target := range currentNode.OutEdges {
     9    152  246:             if number[target] == unvisited {
     7    354  247:                     lastid = DFS(target, nodes, number, last, lastid+1)
     .      .  248:             }
     .      .  249:     }
     7     59  250:     last[number[currentNode]] = lastid
     1      1  251:     return lastid

columns meaning:
sample_count time  source_line

(pprof) top5
Total: 1652 samples
     197  11.9%  11.9%      382  23.1% scanblock
     189  11.4%  23.4%     1549  93.8% main.FindLoops
     130   7.9%  31.2%      152   9.2% sweepspan
     104   6.3%  37.5%      896  54.2% runtime.mallocgc
      98   5.9%  43.5%      100   6.1% flushptrbuf

Now the program is spending most of its time allocating memory and garbage collecting (runtime.mallocgc, which both
allocates and runs periodic garbage collections, accounts for 54.2% of the time). To find out why the garbage collector
is running so much, we have to find out what is allocating memory.
*/

const letter = `hello {{.Name}}!`

var globalTemplate = template.Must(template.New(`hello {{.Name}}!`).Parse(letter))

type dataStruct struct {
	Name string
}

func parseTemplate(data dataStruct) (string, error) {
	t := template.Must(template.New("letter").Parse(letter))

	var result bytes.Buffer
	err := t.Execute(&result, data)
	return result.String(), err
}

func executeTemplate(t *template.Template, data dataStruct) (string, error) {
	var result bytes.Buffer
	err := t.Execute(&result, data)
	return result.String(), err
}

func Test_ParseTemplate(t *testing.T) {
	expected := `hello yumcoder!`
	result, err := parseTemplate(dataStruct{"yumcoder"})
	if err != nil {
		t.Errorf("executing template: %s", err)
	}

	if expected != result {
		t.Error("expected = ", expected, ", result = ", result)
	}
}

func Test_ExecuteTemplate(t *testing.T) {
	expected := `hello yumcoder!`
	result, err := executeTemplate(globalTemplate, dataStruct{"yumcoder"})
	if err != nil {
		t.Errorf("executing template: %s", err)
	}

	if expected != result {
		t.Error("expected = ", expected, ", result = ", result)
	}
}

// 50000	     24724 ns/op	    9896 B/op	      80 allocs/op
func Benchmark_ParseTemplate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseTemplate(dataStruct{"yumcoder"})
	}
}

// 500000	      2028 ns/op	     504 B/op	      12 allocs/op
func Benchmark_ExecuteTemplate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		executeTemplate(globalTemplate, dataStruct{"yumcoder"})
	}
}
