package popcountbottom

import (
	"fmt"
	"gopl.io/ch2/popcount"
	"testing"
)

// -- Test --
func TestPopCount(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(7), 3},
		{uint64(256), 1},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("PopCount(%q)", test.x)
		if got := PopCount(test.x); got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountBottom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}
