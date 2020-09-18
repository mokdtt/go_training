package main

import (
	"fmt"
	"testing"
)

// -- Test --
func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"123", "123"},
		{"12345", "12,345"},
		{"12345678", "12,345,678"},
		{"12345.678", "12,345.678"},
		{"123456.789", "123,456.789"},
		{"1.23456", "1.23456"},
		{"0.12345678", "0.12345678"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("Comma(%q)", test.s)
		if got := CommaFloat(test.s); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
