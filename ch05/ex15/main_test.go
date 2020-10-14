package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		args []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{0, -1, -2, -3, -4, -5}, 0},
		{[]int{-2, -1, 0, 1, 2}, 2},
		{[]int{}, 0}, //errorだけど0返ってくる
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max(%q)", test.args)
		got, _ := max(test.args...)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		args []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{0, -1, -2, -3, -4, -5}, -5},
		{[]int{-2, -1, 0, 1, 2}, -2},
		{[]int{}, 0}, //errorだけど0返ってくる
	}

	for _, test := range tests {
		descr := fmt.Sprintf("min(%q)", test.args)
		got, _ := min(test.args...)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
