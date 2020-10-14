package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	sep := ","
	fmt.Println(join(sep, s...))
	s2 := []string{"あ", "い", "う"}
	sep2 := "あ"
	fmt.Println(join(sep2, s2...))
}

func join(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}
	s := ""
	for i, aa := range a {
		if i == 0 {
			s += aa
			continue
		}
		s += sep + aa
	}
	return s
}
