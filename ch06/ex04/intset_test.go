package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestElems(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 70}},
	}

	for _, test := range tests {
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		descr := fmt.Sprintf("*IntSet(%v).Elems()", s.String())
		want := test.input
		got := s.Elems()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
