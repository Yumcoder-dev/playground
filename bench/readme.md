# file template
typical example of a go file format

```
// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// file level description write here
package playground

import "fmt"

func main() {
    fmt.Printf("yumcoder")
}
```

# optimization
Optimizations are usually application-specific, but here are some common suggestions.
for more info see [dave.cheney.net](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

* Combine objects into larger objects. For example, replace *bytes.Buffer struct member with bytes.Buffer (you can
preallocate buffer for writing by calling bytes.Buffer.Grow later). This will reduce number of memory allocations
(faster) and also reduce pressure on garbage collector (faster garbage collections).

* Local variables that escape from their declaration scope get promoted into heap allocations. Compiler generally
can't prove that several variables have the same life time, so it allocates each such variable separately.
So you can use the above advise for local variables as well. For example, replace:
    ```
    for k, v := range m {
        k, v := k, v   // copy for capturing by the goroutine
        go func() {
            // use k and v
        }()
    }
    ```
    with:
    ```
    for k, v := range m {
        x := struct{ k, v string }{k, v}   // copy for capturing by the goroutine
        go func() {
            // use x.k and x.v
        }()
    }
    ```
    This replaces two memory allocations with a single allocation. However, this optimization usually negatively affects
    code readability, so use it reasonably.

* A special case of allocation combining is slice array preallocation. If you know a typical size of the slice,
you can preallocate a backing array for it as follows:
    ```
    type X struct {
        buf      []byte
        bufArray [16]byte // Buf usually does not grow beyond 16 bytes.
    }

    func MakeX() *X {
        x := &X{}
        // Preinitialize buf with the backing array.
        x.buf = x.bufArray[:0]
        return x
    }
    ```
* If possible use smaller data types. For example, use int8 instead of int.

* Objects that do not contain any pointers (note that strings, slices, maps and chans contain implicit pointers),
are not scanned by garbage collector. For example, a 1GB byte slice virtually does not affect garbage collection
time. So if you remove pointers from actively used objects, it can positively impact garbage collection time.
Some possibilities are: replace pointers with indices, split object into two parts one of which does not contain pointers.

* Use freelists to reuse transient objects and reduce number of allocations. Standard library contains sync.Pool
type that allows to reuse the same object several times in between garbage collections. However, be aware that,
as any manual memory management scheme, incorrect use of sync.Pool can lead to use-after-free bugs.

