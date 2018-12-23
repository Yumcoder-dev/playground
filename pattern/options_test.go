// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package pattern

import (
	"fmt"
)

func Host(host string) func(*Server) {
	return func(s *Server) {
		s.Host = host
	}
}

func Port(port int) func(*Server) {
	return func(s *Server) {
		s.Port = port
	}
}

type Server struct {
	Host string
	Port int
}

func NewServer(opts ...func(*Server)) *Server {
	s := &Server{}

	// call option functions on instance to set options on it
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func Example_FunctionalOption() {
	s1 := NewServer(Host("127.0.0.1"), Port(8080))
	fmt.Printf("server host: %s, port: %d\n", s1.Host, s1.Port)

	s2 := NewServer(Host("127.0.0.1"))
	fmt.Printf("server host: %s\n", s2.Host)

	// Output:
	// server host: 127.0.0.1, port: 8080
	// server host: 127.0.0.1
}
