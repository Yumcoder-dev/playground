// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"crypto/rand"
	"crypto/sha1"
	"testing"
)

// https://blog.golang.org/go-slices-usage-and-internals
//
// array [count]type
// An array type definition specifies a length and an element type. For example, the type [4]int represents an array of
// four integers. An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types).
//
// slice []type
// Arrays have their place, but they're a bit inflexible, so you don't see them too often in Go code. Slices, though,
// are everywhere. They build on arrays to provide great power and convenience.
// The type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a
// slice type has no specified length.

func Test_DelAndLen(t *testing.T) {
	s := make([]int, 0, 10)
	t.Log(len(s))
	t.Log(cap(s))
	s = append(s, 10)
	t.Log(len(s))
	t.Log(cap(s))
	s = s[:0]
	t.Log(len(s)) // 0
	t.Log(cap(s)) // 10

	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)
	i := 1
	s = append(s[:i], s[i+1:]...)
	t.Log(len(s)) // 2
	t.Log(cap(s)) // 10
	t.Log(s)
}

func Test_Initial_Byte(t *testing.T) {
	tmp_encrypted_answer := make([]byte, 10)
	t.Log(tmp_encrypted_answer)
}

func Test_Extend_Byte(t *testing.T) {
	handshake := make([]byte, 16)
	authKey := make([]byte, 256)

	rand.Read(handshake)
	rand.Read(authKey)
	t.Log(handshake)
	t.Log(authKey)
	authKeyAuxHash := make([]byte, len(handshake))
	t.Log(authKeyAuxHash)
	copy(authKeyAuxHash, handshake)
	t.Log(authKeyAuxHash)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	t.Log(authKeyAuxHash)
	sha1_d := sha1.Sum(authKey)
	t.Log(sha1_d)
	authKeyAuxHash = append(authKeyAuxHash, sha1_d[:]...)
	t.Log(authKeyAuxHash)

	sha1_e := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	t.Log(sha1_e)
	authKeyAuxHash = append(authKeyAuxHash, sha1_e[:]...)
	t.Log(authKeyAuxHash)

}

func Test_Mack(t *testing.T) {
	b := make([]byte, 10)
	t.Log(b)
	b[8] = 0x4
	b = make([]byte, 15)
	t.Log(b)
}

func Test_append(t *testing.T) {
	type S struct {
		a string
	}
	s1 := make([]*S, 0)
	s1 = append(s1, &S{})

	var s2 []*S
	if len(s1) > 2{
		s2 = append(s2, &S{})
	}
	t.Log("end...")
}
