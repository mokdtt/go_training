// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test of echo command.  Run with: go test gopl.io/ch11/echo

//!+
package main

import (
	"fmt"
	"testing"
)

func TestAddPrefix(t *testing.T) {
	var tests = []struct {
		url  string
		want string
	}{
		{"http://sample.com", "http://sample.com"},
		{"example.com", "http://example.com"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("AddPrefix(%q)", test.url)
		// 判定
		if got := AddPrefix(test.url); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

//!-
