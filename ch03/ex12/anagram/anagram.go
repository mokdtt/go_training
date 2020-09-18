package anagram

import (
	"reflect"
)

func CheckAnagram(s1 string, s2 string) bool {
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	for _, s := range s1 {
		m1[s] += 1
	}
	for _, s := range s2 {
		m2[s] += 1
	}

	return reflect.DeepEqual(m1, m2)
}
