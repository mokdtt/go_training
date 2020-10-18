// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{-1, -2}, 0},
		{[]int{-1, -2, 1}, 1},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("s.Len()")
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		if got := s.Len(); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input  []int
		remove int
		want   []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("s.Len()")
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		var swant IntSet
		for _, v := range test.want {
			swant.Add(v)
		}
		s.Remove(test.remove)
		got := s.String()
		want := swant.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("s.Len()")
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		s.Clear()
		got := s.String()
		want := "{}"
		if got != "{}" {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("s.Copy()")
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		scopy := s.Copy()
		got := scopy.String()
		want := s.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
