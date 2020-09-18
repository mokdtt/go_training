package anagram

import (
	"fmt"
	"testing"
)

// -- Test --
func TestCheckAnagram(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"abc", "bca", true},
		{"a", "b", false},
		{"1234abcd", "abcd1234", true},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("CheckAnagram(%q, %q)", test.s1, test.s2)
		if got := CheckAnagram(test.s1, test.s2); got != test.want {
			t.Errorf("%s = %t, want %t", descr, got, test.want)
		}
	}
}
