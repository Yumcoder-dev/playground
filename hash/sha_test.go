// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package hash

import (
	"crypto/sha1"
	"testing"
)

func Test_sha(t *testing.T) {
	s := sha1.New()
	t.Log(s.Sum([]byte("yumd-pc")))
}
