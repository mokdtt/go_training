// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test of countLilnes command.

//!+
package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestCountLines(t *testing.T) {
	var tests = []struct {
		filename string
		want     map[string]int
	}{
		{"test1.txt", map[string]int{"a": 2, "b": 2}},
		{"test2.txt", map[string]int{"1": 3, "2": 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countLines(%q)", test.filename)
		counts := make(map[string]int)          //文字の出現回数
		flgs := make(map[string]map[string]int) //["a"]["in1.txt"]のような二重構造
		f, err := os.Open(test.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		if err := countLines(f, test.filename, counts, flgs); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		// 判定
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("%s = %q, want %q", descr, counts, test.want)
		}
	}
}

//!-
