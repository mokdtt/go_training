package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "a", "a", "b", "b", "c"}
	s = rmNeighborDup(s)
	fmt.Println(s)
}

func rmNeighborDup(strings []string) []string {
	if len(strings) <= 0 {
		return strings
	}
	i := 0
	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}
