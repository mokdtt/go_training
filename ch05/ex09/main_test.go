package main

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	toaaa := func(s string) string { return "aaa" }
	addあ := func(s string) string { return s + "あ" }
	var tests = []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"$Hello world", toaaa, "aaa world"},
		{"$Hello world $こんにちは 世界", toaaa, "aaa world aaa 世界"},
		{"Hello world", toaaa, "Hello world"},
		{"$Hello world", addあ, "Helloあ world"},
		{"$Hello world $こんにちは 世界", addあ, "Helloあ world こんにちはあ 世界"},
		{"", addあ, ""},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("expand(%q, f)", test.s)
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
