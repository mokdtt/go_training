package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestIsPalindromeInt(t *testing.T) {
	tests := []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3, 3, 2, 1}, true},
		{[]int{100, 2, 3, 2, 1}, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IsPalindrome(%v)", test.s)
		got := IsPalindrome(sort.IntSlice(test.s))
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		s    []string
		want bool
	}{
		{[]string{}, true},
		{[]string{"aa", "bb", "cc", "bb", "aa"}, true},
		{[]string{"aa", "bb", "cc", "cc", "bb", "aa"}, true},
		{[]string{"aaa", "bb", "cc", "bb", "aa"}, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IsPalindrome(%v)", test.s)
		got := IsPalindrome(sort.StringSlice(test.s))
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
