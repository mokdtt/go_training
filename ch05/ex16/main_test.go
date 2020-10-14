package main

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	var tests = []struct {
		a    []string
		sep  string
		want string
	}{
		{[]string{""}, ",", ""},
		{[]string{"one"}, ",", "one"},
		{[]string{"one", "two", "three"}, " ", "one two three"},
		{[]string{"a", "b", "c"}, "", "abc"},
		{[]string{"1", "2", "3"}, " ", "1 2 3"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("join(%s, %q)", test.sep, test.a)
		got := join(test.sep, test.a...)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
