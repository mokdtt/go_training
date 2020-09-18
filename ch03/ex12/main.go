package main

import (
	"fmt"
	"go_training/ch03/ex12/anagram"
)

func main() {
	s1 := "abcdef"
	s2 := "defabc"
	fmt.Println(anagram.CheckAnagram(s1, s2))
}
