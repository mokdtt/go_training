package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	toaaa := func(s string) string {
		return "aaa"
	}
	s := "$hello world, $こんにちは 世界"
	fmt.Println(s)
	s = expand(s, toaaa)
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		ss := input.Text()
		if strings.HasPrefix(ss, "$") {
			s = strings.ReplaceAll(s, ss, f(ss[1:]))
		}
	}
	return s
}
