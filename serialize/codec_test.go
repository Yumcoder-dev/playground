// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package serialize

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"math/rand"
	"testing"
	"time"
	"yumcoder.com/yumd/server/core/proto/schema"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func func_msgpack() {
	salt := &schema.TLFutureSalt{Data2:&schema.FutureSalt_Data{
		Salt:rand.Int63(),
		ValidSince:int32(time.Now().Unix()),
	}}
	salts := []*schema.TLFutureSalt{salt}
	b, err := msgpack.Marshal(salts)
	if err != nil {
		panic(err)
	}
	fmt.Println("func_msgpack", len(b))

	var item []*schema.TLFutureSalt
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	//fmt.Println(item)
}

func func_json() {
	salt := &schema.TLFutureSalt{Data2:&schema.FutureSalt_Data{
		Salt:rand.Int63(),
		ValidSince:int32(time.Now().Unix()),
	}}
	salts := []*schema.TLFutureSalt{salt}
	b, err := json.Marshal(salts)
	if err != nil {
		panic(err)
	}

	fmt.Println("func_json", len(b))

	var item []*schema.TLFutureSalt
	err = json.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	//fmt.Println(item)
}

func func_GobBase64() {
	salt := &schema.TLFutureSalt{Data2:&schema.FutureSalt_Data{
		Salt:rand.Int63(),
		ValidSince:int32(time.Now().Unix()),
	}}
	salts := []*schema.TLFutureSalt{salt}
	b := toGobBase64(salts)

	fmt.Println("func_GobBase64", len(b))


	var item []*schema.TLFutureSalt
	item = fromGobBase64(b)
	_ = item
	//fmt.Println(item)
}
func TestName(t *testing.T) {
	func_GobBase64()
}

func Benchmark_serialization(b *testing.B) {
	benchTable := []struct {
		name string
		f func()
	}{
		{"msgpack", func() {func_msgpack()}},
		{"json", func() {func_json()}},
		{"GobBase64", func() {func_GobBase64()}},
	}
	for _, v := range benchTable{
		b.Run(v.name, func(b *testing.B) {
			b.ReportAllocs()
			for i:= 0;i<1000;i++ {
				v.f()
			}
		})
	}
}

func toGobBase64(m []*schema.TLFutureSalt) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil { fmt.Println(`failed gob Encode`, err) }
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func fromGobBase64(str string) []*schema.TLFutureSalt {
	m := make([]*schema.TLFutureSalt, 0)
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil { fmt.Println(`failed base64 Decode`, err); }
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil { fmt.Println(`failed gob Decode`, err); }
	return m
}