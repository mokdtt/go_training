package main

import (
	"fmt"
	"testing"
)

// -- Test --
func TestReverse(t *testing.T) {
	var tests = []struct {
		a    [6]int
		want [6]int
	}{
		{[6]int{1, 2, 2, 3, 3, 4}, [6]int{4, 3, 3, 2, 2, 1}},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", test.a)
		reverse(&test.a)
		got := test.a
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
