package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{"", map[rune]int{}, [utf8.UTFMax + 1]int{}, 0},
		{"aa", map[rune]int{'a': 2}, [utf8.UTFMax + 1]int{1: 2}, 0},
		{"ああいい", map[rune]int{'あ': 2, 'い': 2}, [utf8.UTFMax + 1]int{3: 4}, 0},
		{"😎😎😎", map[rune]int{'😎': 3}, [utf8.UTFMax + 1]int{4: 3}, 0},
	}

	for _, test := range tests {
		counts, utflen, invalid := CharCount(bufio.NewReader(strings.NewReader(test.input)))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("counts: %x, want %x", counts, test.counts)
		}

		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("utflen: %x, want %x", utflen, test.utflen)
		}

		if invalid != test.invalid {
			t.Errorf("invalid: %x, want %x", invalid, test.invalid)
		}
	}
}
