// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package err

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func errEmbedFunc(i string) {
	errorHandling := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err := strconv.Atoi(i)
	errorHandling(err)
	_, err = strconv.Atoi(i)
	errorHandling(err)
}

func errInline(i string) {
	_, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	_, err = strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
}

func Benchmark_CheckError(b *testing.B) {
	benchTable := []struct {
		name  string
		input string
		f     func(i string)
	}{
		{"errEmbedFunc", "10", errEmbedFunc},
		{"errInline", "10", errInline},
	}

	for _, v := range benchTable {
		b.Run(v.name, func(b *testing.B) {
			b.ReportAllocs()
			v.f(v.input)
		})
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func funcHandleResource(i int) (result string, err error) {
	defer func() {
		// resource management when err occurred
		if err != nil {
			fmt.Println("err:", err)
			fmt.Println("release resources...")
		}
	}()

	if i%2 == 0 {
		// throw err
		return "", errors.New("throw err")
	}

	// normal call
	return "result", nil
}

func Test_Err(t *testing.T) {
	r, _ := funcHandleResource(1)
	t.Logf("funcHandleResource(1): %s", r)

	r, _ = funcHandleResource(2)
	t.Logf("funcHandleResource(2): %s", r)
}
