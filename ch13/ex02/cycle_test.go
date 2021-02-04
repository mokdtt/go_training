package cycle

import (
	"bytes"
	"testing"
)

func TestIsCycle(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1 := make(chan int)

	type mystring string

	var iface1 interface{} = &one

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false}, // different values
		{"foo", false},
		{mystring("foo"), false}, // different types
		// slices
		{[]string{"foo"}, false},
		// slice cycles
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
		// functions
		{(func())(nil), false},
		// arrays
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		// interfaces
		{&iface1, false},
	} {
		if IsCyclic(test.x) != test.want {
			t.Errorf("IsCyclic(%v) = %t",
				test.x, !test.want)
		}
	}
}
