// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package datatype

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"github.com/gogo/protobuf/proto"
// 	"testing"
// 	"yumcoder.com/yumd/server/core/test2"
// )

// func Test_Size(t *testing.T) {
// 	b := make([]byte, 500)
// 	rand.Read(b)
// 	m := &test2.PingRequest{
// 		Message: string(b),
// 	}
// 	t.Log(m)
// 	buf, _ := proto.Marshal(m)
// 	t.Log(len(buf))
// 	t.Log(len(base64.StdEncoding.EncodeToString(buf)))
// }