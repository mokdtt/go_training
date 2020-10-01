package main

import (
	"bytes"
	"fmt"
	"testing"
)

// -- Test --
func TestRotate(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"Hello   世界", "Hello 世界"},
		{"Hello   世界   world", "Hello 世界 world"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("compress(%v)", test.s)
		got := compress([]byte(test.s))
		want := []byte(test.want)
		if !bytes.Equal(got, want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
