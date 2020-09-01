package weightconv

import (
	"fmt"
	"testing"
)

func TestPToK(t *testing.T) {
	var tests = []struct {
		p    Pond
		want Kilogram
	}{
		{0.0, 0.0},
		{2.2046, 1.0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PToK(%s)", test.p)
		got := PToK(test.p)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestKToP(t *testing.T) {
	var tests = []struct {
		k    Kilogram
		want Pond
	}{
		{0.0, 0.0},
		{1.0, 2.2046},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToP(%s)", test.k)
		got := KToP(test.k)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
