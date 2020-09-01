package lengthconv

import (
	"fmt"
	"testing"
)

func TestFToM(t *testing.T) {
	var tests = []struct {
		f    Feet
		want Meter
	}{
		{0.0, 0.0},
		{3.2808, 1.0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("FToM(%s)", test.f)
		got := FToM(test.f)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestMToF(t *testing.T) {
	var tests = []struct {
		m    Meter
		want Feet
	}{
		{0.0, 0.0},
		{1.0, 3.2808},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("MToF(%s)", test.m)
		got := MToF(test.m)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
