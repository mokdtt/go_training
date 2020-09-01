// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test of echo command.  Run with: go test gopl.io/ch11/echo

//!+
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"one"}, "0 one\n"},
		{[]string{"one", "two", "three"}, "0 one\n1 two\n2 three\n"},
		{[]string{"a", "b", "c"}, "0 a\n1 b\n2 c\n"},
		{[]string{"1", "2", "3"}, "0 1\n1 2\n2 3\n"},
	}

	for _, test := range tests {
		// ファイル名を追加
		descr := fmt.Sprintf("echo(%q)", test.args)

		// 自作package echoを利用
		out = new(bytes.Buffer) // captured output
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		// 判定
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

//!-
