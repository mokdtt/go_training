// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test of echo command.  Run with: go test gopl.io/ch11/echo

//!+
package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"one"}, "one\n"},
		{[]string{"one", "two", "three"}, "one two three\n"},
		{[]string{"a", "b", "c"}, "a b c\n"},
		{[]string{"1", "2", "3"}, "1 2 3\n"},
	}

	for _, test := range tests {
		// ファイル名を追加
		input := append(os.Args[0:1], test.args...)
		want := os.Args[0] + " " + test.want

		descr := fmt.Sprintf("echo(%q)", input)

		// 自作package echoを利用
		out = new(bytes.Buffer) // captured output
		if err := echo(input); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		// 判定
		got := out.(*bytes.Buffer).String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

//!-
