package main

import (
	"fmt"
	"reflect"
	"testing"
)

// -- Test --
func TestRotate(t *testing.T) {
	var tests = []struct {
		s    []string
		want []string
	}{
		{[]string{"a", "a", "a", "b", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "a", "a", "a"}, []string{"a"}},
		{[]string{"a", "a", "b", "a", "a"}, []string{"a", "b", "a"}},
		{[]string{}, []string{}},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("rmNeighborDup(%v)", test.s)
		got := rmNeighborDup(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
