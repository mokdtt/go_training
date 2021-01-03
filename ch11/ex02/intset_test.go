package intset

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		ints []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 1, 2}},
		{[]int{0, 1, 2, 3, 1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		intset := IntSet{}
		mapintset := MapIntSet{}
		for _, n := range test.ints {
			intset.Add(n)
			mapintset.Add(n)
		}
		newmap := intset2map(intset)
		if !reflect.DeepEqual(newmap, mapintset.words) {
			t.Errorf("intset: %v != mapintset: %v", newmap, mapintset.words)
		}
	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		ints1 []int
		ints2 []int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{1, 1, 2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		intset1 := IntSet{}
		intset2 := IntSet{}
		mapintset1 := MapIntSet{}
		mapintset2 := MapIntSet{}
		for _, n := range test.ints1 {
			intset1.Add(n)
			mapintset1.Add(n)
		}
		for _, n := range test.ints2 {
			intset2.Add(n)
			mapintset2.Add(n)
		}
		intset1.UnionWith(&intset2)
		mapintset1.UnionWith(&mapintset2)
		newmap := intset2map(intset1)
		if !reflect.DeepEqual(newmap, mapintset1.words) {
			t.Errorf("intset: %v != mapintset: %v", newmap, mapintset1.words)
		}
	}
}

func intset2map(ints IntSet) map[int]bool {
	newmap := map[int]bool{}
	for i, word := range ints.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				newmap[64*i+j] = true
			}
		}
	}
	return newmap
}
