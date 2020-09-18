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
	}
	for _, test := range tests {
		descr := fmt.Sprintf("Comma(%q)", test.s)
		if got := Comma(test.s); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

// -- Benchmarks --

func BenchmarkComma9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma("123456789")
	}
}

func BenchmarkCommaOrig9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommaOrig("123456789")
	}
}

func BenchmarkComma60(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma("123456789999999999999999999999999999999999999999999999999999")
	}
}

func BenchmarkCommaOrig60(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommaOrig("123456789999999999999999999999999999999999999999999999999999")
	}
}
