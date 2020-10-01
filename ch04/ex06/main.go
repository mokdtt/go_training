package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Hello,    世界      world"
	r := []byte(s)
	fmt.Println(string(r))
	r = compress(r)
	fmt.Println(string(r))
}

func compress(bs []byte) []byte {
	if len(bs) == 0 {
		return bs
	}
	loc := 0
	var rprev rune
	for i := 0; i < len(bs); {
		r, size := utf8.DecodeRune(bs[loc:])
		if unicode.IsSpace(r) && unicode.IsSpace(rprev) {
			copy(bs[loc:], bs[loc+size:])
		} else {
			loc += size
		}
		i += size
		rprev = r
	}
	return bs[:loc]
}
