package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"hello world", 2},
		{"hello world　こんにちは", 3},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(*WordCounter).Write(%q)", test.input)
		var c WordCounter
		c.Write([]byte(test.input))

		got := int(c)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"hello world", 1},
		{`aaa
		bbb
		ccc`, 3},
		{`こんにちは
		世界`, 2},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(*LineCounter).Write(%q)", test.input)
		var c LineCounter
		c.Write([]byte(test.input))

		got := int(c)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
