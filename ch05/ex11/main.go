package main

import (
	"fmt"
	"os"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func isAcyclic(done []string, target string, m map[string][]string) bool {
	contains := func(s []string, t string) bool {
		for _, v := range s {
			if v == t {
				return true
			}
		}
		return false
	}
	for _, item := range done {
		if contains(m[item], target) {
			return false
		}
	}
	return true
}

//acyclicでないことを示すには，あるitemをorderに追加する前に
//今まで追加された中に戻るpathがないか確認する必要がある
func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				if ok := isAcyclic(order, item, m); !ok {
					return fmt.Errorf("not acyclic graph\n")
				}
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}
