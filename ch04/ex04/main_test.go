package main

import (
	"fmt"
	"reflect"
	"testing"
)

// -- Test --
func TestRotate(t *testing.T) {
	var tests = []struct {
		a    []int
		n    int
		want []int
	}{
		{[]int{1, 2, 2, 3, 3, 4}, 1, []int{2, 2, 3, 3, 4, 1}},
		{[]int{2, 3, 4, 5}, 2, []int{4, 5, 2, 3}},
		{[]int{2, 3, 4, 5}, 5000, []int{2, 3, 4, 5}},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("rotate(%v, %v)", test.a, test.n)
		rotate(test.a, test.n)
		got := test.a
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
