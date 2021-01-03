package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"", " ", 1},
		{"a:b:c", ":", 3},
		{"a b c d", " ", 4},
		{"あいあいあ", "い", 3},
		{"あ亜あ亜あ", "亜", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
