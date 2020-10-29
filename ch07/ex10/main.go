package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	limit := s.Len() / 2
	for i := 0; i < limit; i++ {
		j := s.Len() - i - 1
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {
	s1 := []int{1, 2, 3, 2, 1}
	fmt.Println(s1, IsPalindrome(sort.IntSlice(s1)))
	s2 := []int{1, 2, 3, 2, 2}
	fmt.Println(s2, IsPalindrome(sort.IntSlice(s2)))
}
