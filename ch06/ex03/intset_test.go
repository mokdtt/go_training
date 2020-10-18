package intset

import (
	"fmt"
	"testing"
)

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 2, 70}, []int{2, 3, 4}, []int{2}},
		{[]int{1, 2, 3}, []int{2, 3, 4, 70}, []int{2, 3}},
		{[]int{1, 2, 3, 100}, []int{2, 3, 4, 100}, []int{2, 3, 100}},
	}

	for _, test := range tests {
		var s1 IntSet
		for _, v := range test.input1 {
			s1.Add(v)
		}
		var s2 IntSet
		for _, v := range test.input2 {
			s2.Add(v)
		}
		descr := fmt.Sprintf("*IntSet(%v).IntersectWith(%v)", s1.String(), s2.String())
		var swant IntSet
		for _, v := range test.want {
			swant.Add(v)
		}
		s1.IntersectWith(&s2)
		got := s1.String()
		want := swant.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1}},
		{[]int{1, 2, 70}, []int{2, 3, 4}, []int{1, 70}},
		{[]int{1, 2, 3}, []int{2, 3, 4, 70}, []int{1}},
		{[]int{1, 2, 3, 100}, []int{2, 3, 4, 100}, []int{1}},
	}

	for _, test := range tests {
		var s1 IntSet
		for _, v := range test.input1 {
			s1.Add(v)
		}
		var s2 IntSet
		for _, v := range test.input2 {
			s2.Add(v)
		}
		descr := fmt.Sprintf("*IntSet(%v).DifferenceWith(%v)", s1.String(), s2.String())
		var swant IntSet
		for _, v := range test.want {
			swant.Add(v)
		}
		s1.DifferenceWith(&s2)
		got := s1.String()
		want := swant.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1, 4}},
		{[]int{1, 2, 70}, []int{2, 3, 4}, []int{1, 3, 4, 70}},
		{[]int{1, 2, 3}, []int{2, 3, 4, 70}, []int{1, 4, 70}},
		{[]int{1, 2, 3, 100}, []int{2, 3, 4, 100}, []int{1, 4}},
	}

	for _, test := range tests {
		var s1 IntSet
		for _, v := range test.input1 {
			s1.Add(v)
		}
		var s2 IntSet
		for _, v := range test.input2 {
			s2.Add(v)
		}
		descr := fmt.Sprintf("*IntSet(%v).SymmetricDifference(%v)", s1.String(), s2.String())
		var swant IntSet
		for _, v := range test.want {
			swant.Add(v)
		}
		s1.SymmetricDifference(&s2)
		got := s1.String()
		want := swant.String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
