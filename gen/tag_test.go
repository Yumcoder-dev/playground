// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
package gen

import (
	"reflect"
	"testing"
)

func Test_getTag(t *testing.T) {
	type User struct {
		Name  string `mytag:"MyName"`
		Email string `mytag:"MyEmail" yourtag:"123"`
	}

	u := User{"Bob", "bob@mycompany.com"}
	typ := reflect.TypeOf(u)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := typ.FieldByName(fieldName)
		if !found {
			continue
		}
		t.Logf("\nField: User.%s\n", fieldName)
		t.Logf("\tWhole tag value : %q\n", field.Tag)
		t.Logf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))

		// In Go, struct tags are defined as string literals. You cannot
		// directly set the value of a struct tag to be an integer or any
		// other type besides a string. Struct tags are primarily used for
		// metadata, such as specifying field names during serialization or
		// deserialization, and they are intended to be human-readable and
		// easily parsed as strings.
	}
}
