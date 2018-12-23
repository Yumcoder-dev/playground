// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"yumcoder.com/yumd/server/core/crypto2"
)

func Test_Byte(t *testing.T) {
	//a := uint32(7)
	//buf := make([]byte, 4)
	//binary.LittleEndian.PutUint32(buf, a)
	//t.Log(buf)

	buf := new(bytes.Buffer)
	a := int32(-7)
	binary.Write(buf, binary.LittleEndian, a)
	t.Log(buf.Bytes())

	readBuf := buf.Bytes()[:]
	t.Log((uint32(readBuf[2]) << 24) | (uint32(readBuf[1]) << 16) | (uint32(readBuf[0]) << 8) | (uint32(readBuf[0])))
	t.Log(
		(int32(readBuf[3]) << 24) |
			(int32(readBuf[2]) << 16) |
			(int32(readBuf[1]) << 8) |
			(int32(readBuf[0])))

	uVal := ((readBuf[3] & 0xff) << 24) |
		((readBuf[2] & 0xff) << 16) |
		((readBuf[1] & 0xff) << 8) |
		(readBuf[0] & 0xff)
	t.Log(int32(uVal))
}

func Test_Byte2(t *testing.T) {
	a := uint32(7)
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, a)
	t.Log(buf[0:2])

	t.Log(byte(255) >> 7)
	t.Log(byte(1) >> 7)
	t.Log(byte(1) << 2)
}

func Test_AppCoded_e(t *testing.T) {
	randBytes := []byte{
		53, 143, 121, 153, 232, 30, 16, 201, 48, 56, 211, 90, 56, 88, 225, 201, 151, 51, 123, 225, 54, 60, 51, 81, 163, 191, 141, 202, 157, 208, 190,
		208, 90, 80, 95, 4, 42, 232, 242, 22, 210, 139, 244, 8, 238, 111, 121, 140, 79, 56, 45, 62, 117, 174, 122, 84, 239, 239, 239, 239, 148, 5, 252, 108,
	}

	d, _ := crypto2.NewAesCTR128Encrypt(randBytes[8:40], randBytes[40:56])

	encodeBytes := []byte{
		53, 143, 121, 153, 232, 30, 16, 201, 48, 56, 211, 90, 56, 88, 225, 201, 151, 51, 123, 225, 54, 60, 51, 81, 163, 191, 141, 202, 157, 208, 190,
		208, 90, 80, 95, 4, 42, 232, 242, 22, 210, 139, 244, 8, 238, 111, 121, 140, 79, 56, 45, 62, 117, 174, 122, 84, 161, 32, 225, 254, 9, 90, 143, 78,
	}

	b_0_4 := make([]byte, 4)
	copy(b_0_4, encodeBytes[0:4])
	b_4_60 := make([]byte, 60)
	copy(b_4_60, encodeBytes[4:64])
	d.Encrypt(b_0_4)
	d.Encrypt(b_4_60)
	decodeRes := make([]byte, 64)
	copy(decodeRes[:56], randBytes[:56])
	copy(decodeRes[56:64], b_4_60[52:60])
	fmt.Println("decodeRes:", decodeRes)
	assert.Equal(t, randBytes, decodeRes)

	b_0_64 := make([]byte, 64)
	copy(b_0_64, randBytes)
	d, _ = crypto2.NewAesCTR128Encrypt(randBytes[8:40], randBytes[40:56])
	d.Encrypt(b_0_64)
	encodeRes := make([]byte, 64)
	copy(encodeRes[:56], encodeBytes[:56])
	copy(encodeRes[56:64], b_0_64[56:64])
	fmt.Println("encodeRes:", encodeRes)

	assert.Equal(t, encodeBytes, encodeRes)
}

func Test_Copy(t *testing.T) {
	b := make([]byte, 10)
	t.Log(b)

	b1 := make([]byte, 5)
	for i := 0; i < 5; i++ {
		b1[i] = uint8(i + 1)
	}
	//copy(b, b1)
	//t.Log(b)
	//copy(b[4:], b1)
	//t.Log(b)
	a := 12
	copy(b[10-a:], b1)
	t.Log(b)
}
