package main

import (
	"fmt"
	"testing"
)

// -- Test --
func TestDiffCount(t *testing.T) {
	var tests = []struct {
		b1   byte
		b2   byte
		want int
	}{
		{1, 1, 8},
		{2, 1, 6},
		{123, 132, 0},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("DiffCount(%b, %b)", test.b1, test.b2)
		if got := DiffCount(test.b1, test.b2); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
