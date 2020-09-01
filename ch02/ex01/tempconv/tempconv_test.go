package tempconv

import (
	"fmt"
	"testing"
)

func TestKToC(t *testing.T) {
	var tests = []struct {
		k    Kelvin
		want Celsius
	}{
		{0.0, -273.15},
		{273.15, 0.0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToC(%s)", test.k)
		got := KToC(test.k)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		c    Celsius
		want Kelvin
	}{
		{0.0, 273.15},
		{5.0, 278.15},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToK(%s)", test.c)
		got := CToK(test.c)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
