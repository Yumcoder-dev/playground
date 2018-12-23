// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func Test_Compare(t *testing.T) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	p := P{3, 4, 5, "yumcoder"}
	err := enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	t.Log(len(network.Bytes()))
	t.Log(len(val))
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func Example_basic() {
	// Initialize the encoder and decoder. Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "yumcoder"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	// Output:
	// "yumcoder": {3, 4}
	// "Treehouse": {1782, 1841}

}
