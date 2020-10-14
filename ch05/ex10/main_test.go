package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},

		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},

		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	for i := 0; i < 5; i++ {
		sorted := topoSort(prereqs)
		checkSorted(prereqs, sorted, t)
	}
}

func checkSorted(prereqs map[string][]string, sorted []string, t *testing.T) {
	keys := make(map[string]bool)
	for key := range prereqs {
		keys[key] = true
	}
	for item := range keys {
		for _, v := range prereqs[item] {
			pre := find(sorted, item)
			if pre == -1 {
				t.Errorf("not contains %s", item)
			}
			post := find(sorted, v)
			if post == -1 {
				t.Errorf("not contains %s", v)
			}
			if pre == -1 {
				t.Errorf("not contains %s", item)
			}
			if pre < post {
				t.Errorf("want (%s) > (%s), but got (%s) < (%s)", item, v, v, item)
			}
		}
	}
}

func find(s []string, v string) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}
