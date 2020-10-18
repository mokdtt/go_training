package intset

import (
	"fmt"
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		input []int
		add   []int
		want  []int
	}{
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 80, 200}, []int{1, 2, 3}, []int{1, 2, 3, 80, 200}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{6, 5, 4}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		var s IntSet
		for _, v := range test.input {
			s.Add(v)
		}
		descr := fmt.Sprintf("*IntSet(%v).AddAll(%v)", s.String(), test.add)
		var swant IntSet
		for _, v := range test.want {
			swant.Add(v)
		}
		s.AddAll(test.add...)
		got := s.String()
		want := swant.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
