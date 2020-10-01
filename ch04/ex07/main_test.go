package main

import (
	"bytes"
	"fmt"
	"testing"
)

// -- Test --
func TestReverse(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"Hello, 世界", "界世 ,olleH"},
		{"アイウエオ", "オエウイア"},
		{"     あああ", "あああ     "},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", test.s)
		input := []byte(test.s)
		reverse(input)
		got := input
		want := []byte(test.want)
		if !bytes.Equal(got, want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
