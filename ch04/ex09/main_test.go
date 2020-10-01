package main

import (
	"fmt"
	"reflect"
	"testing"
)

// -- Test --
func TestWordFreq(t *testing.T) {
	var tests = []struct {
		s    string
		want map[string]int
	}{
		{"test1.txt", map[string]int{"good": 3, "thank": 4, "you": 4}},
		{"test2.txt", map[string]int{"こんにちは": 2, "ありがとう": 2, "おはよう": 1}},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("wordfreq(%v)", test.s)
		got := wordfreq(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
